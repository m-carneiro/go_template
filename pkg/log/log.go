package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log interface {
	Level() zapcore.Level
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Sync() error
}

func New(z *zap.SugaredLogger) Log {
	return &log{
		z: z,
	}
}

type log struct {
	z *zap.SugaredLogger
}

func (l *log) Level() zapcore.Level {
	return l.z.Level()
}

func (l *log) Info(args ...interface{}) {
	l.z.Info(args)
}

func (l *log) Warn(args ...interface{}) {
	l.z.Warn(args)
}

func (l *log) Error(args ...interface{}) {
	l.z.Error(args)
}

func (l *log) Fatal(args ...interface{}) {
	l.z.Fatal(args)
}

func (l *log) Sync() error {
	return l.z.Sync()
}
