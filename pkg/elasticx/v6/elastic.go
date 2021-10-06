package v6

import (
	"github.com/lixiaolongcuit/gokit/pkg/elasticx"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

func NewElastic(c *elasticx.Config, log *logrus.Entry) (*elastic.Client, error) {
	optitons := []elastic.ClientOptionFunc{}
	if log != nil {
		optitons = append(optitons, elastic.SetInfoLog(&elasticx.LogWrapper{Level: logrus.DebugLevel, LogEntry: log}))
		optitons = append(optitons, elastic.SetErrorLog(&elasticx.LogWrapper{Level: logrus.ErrorLevel, LogEntry: log}))
		optitons = append(optitons, elastic.SetTraceLog(&elasticx.LogWrapper{Level: logrus.TraceLevel, LogEntry: log}))
	}
	if c != nil {
		if len(c.Urls) > 0 {
			optitons = append(optitons, elastic.SetURL(c.Urls...))
		}
		if c.Username != "" {
			optitons = append(optitons, elastic.SetBasicAuth(c.Username, c.Password))
		}
		if c.DisableSniff {
			optitons = append(optitons, elastic.SetSniff(false))
		}
	}
	return elastic.NewClient(optitons...)
}
