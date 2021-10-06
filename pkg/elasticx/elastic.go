package elasticx

import "github.com/sirupsen/logrus"

type Config struct {
	Urls         []string
	Username     string
	Password     string
	DisableSniff bool
}

type LogWrapper struct {
	Level    logrus.Level
	LogEntry *logrus.Entry
}

func (w *LogWrapper) Printf(format string, v ...interface{}) {
	w.LogEntry.Logf(w.Level, format, v...)
}
