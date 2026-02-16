package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(method string, msg string, args ...interface{})
	Info(method string, msg string, args ...interface{})
	Warn(method string, msg string, args ...interface{})
	Error(method string, msg string, err error, args ...interface{})
	Sync() error
	Layer(layer string) Logger
}

type logger struct {
	logger *zap.Logger
}

func NewLogger(level string) (Logger, error) {
	var lvl zapcore.Level

	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}

	zapCfg := zap.NewProductionConfig()
	zapCfg.Level = zap.NewAtomicLevelAt(lvl)
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLogger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}
	return &logger{
		logger: zapLogger,
	}, nil
}

func (l *logger) Debug(method string, msg string, args ...interface{}) {
	l.logger.Debug(
		msg,
		zap.String("method", method),
		zap.Any("details", args),
	)
}

func (l *logger) Info(method string, msg string, args ...interface{}) {
	l.logger.Info(
		msg,
		zap.String("method", method),
		zap.Any("details", args),
	)
}

func (l *logger) Error(method string, msg string, err error, args ...interface{}) {
	l.logger.Error(
		msg,
		zap.String("method", method),
		zap.Error(err),
		zap.Any("details", args),
	)
}

func (l *logger) Warn(method string, msg string, args ...interface{}) {
	l.logger.Warn(
		msg,
		zap.String("method", method),
		zap.Any("details", args),
	)
}

func (l *logger) Layer(layer string) Logger {
	return &logger{
		logger: l.logger.With(zap.String("layer", layer)),
	}
}

func (l *logger) Sync() error {
	return l.logger.Sync()
}
