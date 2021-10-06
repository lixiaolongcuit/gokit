package console

import "testing"

func TestNewTextLogger(t *testing.T) {
	cfg := &Config{
		Level:     "debug",
		Formatter: "text",
	}
	log, err := NewLogger(cfg)
	if err != nil {
		t.Fatal(err)
	}
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
}

func TestNewJsonLogger(t *testing.T) {
	cfg := &Config{
		Level: "debug",
	}
	log, err := NewLogger(cfg)
	if err != nil {
		t.Fatal(err)
	}
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
}
