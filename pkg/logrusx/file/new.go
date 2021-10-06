//输出到文件的logrus配置
package file

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Level     string
	Formatter string
	File      *lumberjack.Logger
}

func NewLogger(c *Config) (*logrus.Logger, error) {
	log := logrus.New()
	if c.Level == "" {
		c.Level = "debug"
	}
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return nil, err
	}
	log.SetLevel(level)
	if c.Formatter == "" {
		c.Formatter = "json"
	}
	if c.Formatter == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else if c.Formatter == "text" {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	} else {
		return nil, fmt.Errorf("log formatter error: %s", c.Formatter)
	}
	if c.File != nil {
		log.SetOutput(c.File)
	}
	return log, nil
}
