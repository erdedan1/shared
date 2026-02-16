package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(layer string, method string, msg string, args ...interface{})
	Info(layer string, method string, msg string, args ...interface{})
	Warn(layer string, method string, msg string, args ...interface{})
	Error(layer string, method string, msg string, err error, args ...interface{})
	Sync() error
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

func (l *logger) Debug(layer string, method string, msg string, args ...interface{}) {
	l.logger.Debug(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Any("details", args),
	)
}

func (l *logger) Info(layer string, method string, msg string, args ...interface{}) {
	l.logger.Info(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Any("details", args),
	)
}

func (l *logger) Error(layer string, method string, msg string, err error, args ...interface{}) {
	l.logger.Error(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Error(err),
		zap.Any("details", args),
	)
}

func (l *logger) Warn(layer string, method string, msg string, args ...interface{}) {
	l.logger.Warn(
		msg,
		zap.String("layer", layer),
		zap.String("method", method),
		zap.Any("details", args),
	)
}

func (l *logger) With(args ...zap.Field) Logger {
	return &logger{
		logger: l.logger.With(args...),
	}
}

func (l *logger) Sync() error {
	return l.logger.Sync()
}
