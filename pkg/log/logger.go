package log

import (
	"go.uber.org/zap"
)

type (
	LogInfo interface {
		Debug(args ...interface{})
		Info(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
		Panic(args ...interface{})
		Fatal(args ...interface{})
	}

	LogFormat interface {
		Debugf(template string, args ...interface{})
		Infof(template string, args ...interface{})
		Warnf(template string, args ...interface{})
		Errorf(template string, args ...interface{})
		Panicf(template string, args ...interface{})
		Fatalf(template string, args ...interface{})
	}

	LogInfoFormat interface {
		LogInfo
		LogFormat
	}
)

type Logger struct {
	zapSugarLogger *zap.SugaredLogger
}

var std *zap.Logger

func Init() {
	logger, _ := NewZapLogger()
	std = logger.zaplogger
}

func Debug(args ...interface{}) {
	std.Sugar().Debug(args)
}
func (l *Logger) Debug(args ...interface{}) {
	l.zapSugarLogger.Debug(args)
}

func Info(args ...interface{}) {
	std.Sugar().Info(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.zapSugarLogger.Info(args)
}

func Warn(args ...interface{}) {
	std.Sugar().Warn(args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.zapSugarLogger.Warn(args)
}

func Error(args ...interface{}) {
	std.Sugar().Error(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.zapSugarLogger.Error(args)
}

func Panic(args ...interface{}) {
	std.Sugar().Panic(args)
}

func (l *Logger) Panic(args ...interface{}) {
	l.zapSugarLogger.Panic(args)
}

func Fatal(args ...interface{}) {
	std.Sugar().Fatal(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.zapSugarLogger.Fatal(args)
}

func Debugf(template string, args ...interface{}) {
	std.Sugar().Debugf(template, args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.zapSugarLogger.Debugf(template, args)
}

func Infof(template string, args ...interface{}) {
	std.Sugar().Infof(template, args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.zapSugarLogger.Infof(template, args)
}

func Warnf(template string, args ...interface{}) {
	std.Sugar().Warnf(template, args)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.zapSugarLogger.Warnf(template, args)
}

func Errorf(template string, args ...interface{}) {
	std.Sugar().Errorf(template, args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.zapSugarLogger.Errorf(template, args)
}

func Panicf(template string, args ...interface{}) {
	std.Sugar().Panicf(template, args)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.zapSugarLogger.Panicf(template, args)
}

func Fatalf(template string, args ...interface{}) {
	std.Sugar().Fatalf(template, args)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.zapSugarLogger.Fatalf(template, args)
}

func Flush() {
	_ = std.Sync()
}
