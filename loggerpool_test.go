package go_logger

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Error("error:", err)
	}

	config := &LoggerConfigure{
		Name:         "app",
		Path:         cwd,
		Level:        LevelWarn,
		MaxAge:       24,
		RotationTime: 240,
	}

	SetLoggerConfig(LoggerTypeZap, config)

	Debug("test debug")
	Info("test info")
	Warn("test warn")
	Error("test error")

	Debugf("test %s", "debugf")
	Infof("test %s", "infof")
	Warnf("test %s", "warnf")
	Errorf("test %s", "errorf")
}
