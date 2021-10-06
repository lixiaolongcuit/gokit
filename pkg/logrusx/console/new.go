//输出到标准输出的配置
package console

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Level     string
	Formatter string
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
	return log, nil
}
