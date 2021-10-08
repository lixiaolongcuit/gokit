package app

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lixiaolongcuit/gokit/pkg/grpcx"
	"github.com/sirupsen/logrus"
)

type Option func(app *App)

func GrpcServer(server *grpcx.Server) Option {
	return func(app *App) { app.grpcServer = server }
}

func GrpcGatewayServerCreater(creater func() (*http.Server, error)) Option {
	return func(app *App) { app.grpcGatewayServerCreater = creater }
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
	grpcServer               *grpcx.Server
	grpcGatewayServer        *http.Server
	grpcGatewayServerCreater func() (*http.Server, error)
	prometheusServer         *http.Server
	log                      *logrus.Entry
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
	if app.grpcGatewayServerCreater != nil {
		go func(errChan chan<- error) {
			var err error
			app.grpcGatewayServer, err = app.grpcGatewayServerCreater()
			if err != nil {
				errChan <- err
				return
			}
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
		if err := app.grpcServer.Close(); err != nil {
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
