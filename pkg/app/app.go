package app

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcServer struct {
	grpcServer *grpc.Server
	addr       string
	tcpServer  net.Listener
}

func NewGrpcServer(server *grpc.Server, addr string) *grpcServer {
	return &grpcServer{
		grpcServer: server,
		addr:       addr,
	}
}

func (s *grpcServer) ListenAndServe() error {
	tcpServer, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	if err := s.grpcServer.Serve(tcpServer); err != nil {
		return err
	}
	return nil
}

type Option func(app *App)

func GrpcServer(server *grpcServer) Option {
	return func(app *App) { app.grpcServer = server }
}

func GrpcGatewayServer(server *http.Server) Option {
	return func(app *App) { app.grpcGatewayServer = server }
}

func PrometheusServer(server *http.Server) Option {
	return func(app *App) { app.prometheusServer = server }
}

func NewApp(log *logrus.Entry, opts ...Option) *App {
	app := &App{
		log: log,
	}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

type App struct {
	grpcServer        *grpcServer
	grpcGatewayServer *http.Server
	prometheusServer  *http.Server
	log               *logrus.Entry
}

func (app *App) Run() {
	var errChan = make(chan error)
	//启动grpc服务
	if app.grpcServer != nil {
		go func(errChan chan<- error) {
			err := app.grpcServer.ListenAndServe()
			if err != nil {
				errChan <- err
			}
		}(errChan)
	}
	//启动grpc-gateway
	if app.grpcGatewayServer != nil {
		go func(errChan chan<- error) {
			if err := app.grpcGatewayServer.ListenAndServe(); err != nil {
				if err != nil {
					errChan <- err
				}
			}
		}(errChan)
	}
	//启动prometheusHttpServer
	if app.prometheusServer != nil {
		go func(errChan chan<- error) {
			if err := app.prometheusServer.ListenAndServe(); err != nil {
				if err != nil {
					errChan <- err
				}
			}
		}(errChan)
	}
	//优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	var err error
	select {
	case err = <-errChan:
	case <-quit:
	}
	if err != nil {
		app.log.Errorf("server error: %+v", err)
	}
	if app.grpcServer != nil {
		if err := app.grpcServer.tcpServer.Close(); err != nil {
			app.log.Errorf("close grpc server error: %+v", err)
		}
	}
	if app.grpcGatewayServer != nil {
		if err := app.grpcGatewayServer.Close(); err != nil {
			app.log.Errorf("close grpc gateway server error: %+v", err)
		}
	}
	if app.prometheusServer != nil {
		if err := app.prometheusServer.Close(); err != nil {
			app.log.Errorf("close prometheus server error: %+v", err)
		}
	}
}
