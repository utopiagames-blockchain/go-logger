package go_logger

import (
	"log"
	"os"
)

var (
	logger Logger
)

func SetLoggerConfig(loggerType LoggerType, cfg *LoggerConfigure) {
	logger = NewLogger(loggerType, cfg)
}

func GetLogger() Logger {
	if logger == nil {
		logger = NewLogger(LoggerTypeZap, &LoggerConfigure{
			Name:         "app",
			Level:        LevelInfo,
			RotationTime: 24,  // 1 day
			MaxAge:       240, // 10 days
		})
	}
	return logger
}

func initLogger() {
	if logger == nil {
		dir, err := os.Getwd()
		if err != nil {
			log.Panicf("could not get current working directory ")
		}
		cfg := &LoggerConfigure{
			Name:         "app",
			Path:         dir,
			Level:        LevelDebug,
			MaxAge:       24,
			RotationTime: 0,
		}
		logger = newZapLogger(cfg)
	}
}

func Fatal(args ...interface{}) {
	initLogger()
	logger.Fatal(args...)
}

func Error(args ...interface{}) {
	initLogger()
	logger.Error(args...)
}

func Warn(args ...interface{}) {
	initLogger()
	logger.Warn(args...)
}

func Info(args ...interface{}) {
	initLogger()
	logger.Info(args...)
}

func Debug(args ...interface{}) {
	initLogger()
	logger.Debug(args...)
}

func Fatalf(fmt string, args ...interface{}) {
	initLogger()
	logger.Fatalf(fmt, args...)
}

func Errorf(fmt string, args ...interface{}) {
	initLogger()
	logger.Errorf(fmt, args...)
}

func Warnf(fmt string, args ...interface{}) {
	initLogger()
	logger.Warnf(fmt, args...)
}

func Infof(fmt string, args ...interface{}) {
	initLogger()
	logger.Infof(fmt, args...)
}

func Debugf(fmt string, args ...interface{}) {
	initLogger()
	logger.Debugf(fmt, args...)
}
