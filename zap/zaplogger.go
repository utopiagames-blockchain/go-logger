package zap

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	fileLogger    *zap.SugaredLogger
	consoleLogger *zap.SugaredLogger
}

type ZapLoggerConfig struct {
	Name         string
	Path         string
	Level        zapcore.Level
	MaxAge       time.Duration
	RotationTime time.Duration
}

func (z *zapLogger) Fatal(args ...interface{}) {
	z.consoleLogger.Fatal(args...)
	z.fileLogger.Fatal(args...)
}
func (z *zapLogger) Error(args ...interface{}) {
	z.consoleLogger.Error(args...)
	z.fileLogger.Error(args...)
}
func (z *zapLogger) Warn(args ...interface{}) {
	z.consoleLogger.Warn(args...)
	z.fileLogger.Warn(args...)
}
func (z *zapLogger) Info(args ...interface{}) {
	z.consoleLogger.Info(args...)
	z.fileLogger.Info(args...)
}
func (z *zapLogger) Debug(args ...interface{}) {
	z.consoleLogger.Debug(args...)
	z.fileLogger.Debug(args...)
}

func (z *zapLogger) Debugf(fmt string, args ...interface{}) {
	z.consoleLogger.Debugf(fmt, args...)
	z.fileLogger.Debugf(fmt, args...)
}

func (z *zapLogger) Infof(fmt string, args ...interface{}) {
	z.consoleLogger.Infof(fmt, args...)
	z.fileLogger.Infof(fmt, args...)
}

func (z *zapLogger) Warnf(fmt string, args ...interface{}) {
	z.consoleLogger.Warnf(fmt, args...)
	z.fileLogger.Warnf(fmt, args...)
}

func (z *zapLogger) Errorf(fmt string, args ...interface{}) {
	z.consoleLogger.Errorf(fmt, args...)
	z.fileLogger.Errorf(fmt, args...)
}

func (z *zapLogger) Fatalf(fmt string, args ...interface{}) {
	z.consoleLogger.Fatalf(fmt, args...)
	z.fileLogger.Fatalf(fmt, args...)
}

func (z *zapLogger) Named(name string) {
	z.consoleLogger = z.consoleLogger.Named(name)
	z.fileLogger = z.fileLogger.Named(name)
}

func NewZapLogger(c ZapLoggerConfig) *zapLogger {
	logPath := c.Path
	if strings.Trim(logPath, " ") == "" {
		log.Print(errors.New("[Warn] must set logger root path in config.yaml (default: /var/log)"))
		logPath = "/var/log"
	}
	logFile := logPath + "/" + c.Name + "-%Y-%m-%d-%H.log"

	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(c.MaxAge*time.Hour),
		rotatelogs.WithRotationTime(c.RotationTime*time.Hour))

	if err != nil {
		panic(err)
	}

	// file logger definition
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		TimeKey:        "timestamp",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	atom := zap.NewAtomicLevelAt(c.Level)

	fileLoggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(rotator),
		atom)

	fileLogger := zap.New(fileLoggerCore, zap.AddCallerSkip(1))

	// console logger definition
	// config := zap.NewProductionConfig()
	// config.Level = atom

	consoleLoggerCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom)

	consoleLogger := zap.New(consoleLoggerCore, zap.AddCallerSkip(1))

	sConsoleLogger := consoleLogger.Sugar().
		WithOptions(zap.AddCallerSkip(2))
	sFileLogger := fileLogger.Sugar().
		WithOptions(zap.AddCallerSkip(2))

	return &zapLogger{
		consoleLogger: sConsoleLogger,
		fileLogger:    sFileLogger,
	}
}
