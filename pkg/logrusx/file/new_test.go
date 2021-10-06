package file

import (
	"testing"

	"gopkg.in/natefinch/lumberjack.v2"
)

func TestFileLogger(t *testing.T) {
	cfg := &Config{
		Level: "debug",
		File: &lumberjack.Logger{
			Filename:   "/tmp/lumberjack-test.log",
			MaxSize:    5,
			MaxBackups: 2,
			Compress:   true,
		},
	}
	log, err := NewLogger(cfg)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 1000000; i++ {
		log.WithField("num", i).Debug("debug")
	}
}
