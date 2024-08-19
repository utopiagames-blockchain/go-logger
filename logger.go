package go_logger

import (
	"github.com/utopiagames-blockchain/go-logger/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

type LoggerType int
type LogLevel int

const (
	// Logger Type
	LoggerTypeZap LoggerType = iota
)

const (
	// Log Level
	LevelFatal LogLevel = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

type Logger interface {
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})

	Fatalf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Infof(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
}

type LoggerConfigure struct {
	Name         string
	Path         string
	Level        LogLevel
	MaxAge       time.Duration
	RotationTime time.Duration
}

func NewLogger(loggerType LoggerType, cfg *LoggerConfigure) Logger {
	var logger Logger

	if LoggerTypeZap == loggerType {
		logger = newZapLogger(cfg)
	} else {
		log.Panicln("could not find logger type: ", loggerType)
	}

	return logger
}

func newZapLogger(c *LoggerConfigure) Logger {

	var cfg zap.ZapLoggerConfig

	var zapLevel zapcore.Level
	if c.Level == LevelDebug {
		zapLevel = zapcore.DebugLevel
	} else if c.Level == LevelError {
		zapLevel = zapcore.ErrorLevel
	} else if c.Level == LevelInfo {
		zapLevel = zapcore.InfoLevel
	} else if c.Level == LevelWarn {
		zapLevel = zapcore.WarnLevel
	} else if c.Level == LevelFatal {
		zapLevel = zapcore.FatalLevel
	}

	cfg.Name = c.Name
	cfg.Level = zapLevel
	cfg.Path = c.Path
	cfg.MaxAge = c.MaxAge
	cfg.RotationTime = c.RotationTime

	log.Printf("config -> %+v", cfg)

	return zap.NewZapLogger(cfg)
}
