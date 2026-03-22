package logger

import (
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

type LoggerField struct {
	field zap.Field
}

func String(key string, value string) LoggerField {
    return LoggerField{
        field: zap.String(key, value),
    }
}

func Error(err error) LoggerField {
    return LoggerField{
        field: zap.Error(err),
    }
}

func Int(key string, value int) LoggerField {
    return LoggerField{
        field: zap.Int(key, value),
    }
}

func Int64(key string, value int64) LoggerField {
    return LoggerField{
        field: zap.Int64(key, value),
    }
}

func Bool(key string, value bool) LoggerField {
    return LoggerField{
        field: zap.Bool(key, value),
    }
}

func Float64(key string, value float64) LoggerField {
    return LoggerField{
        field: zap.Float64(key, value),
    }
}

func Any(key string, value interface{}) LoggerField {
    return LoggerField{
        field: zap.Any(key, value),
    }
}

func Duration(key string, value time.Duration) LoggerField {
    return LoggerField{
        field: zap.Duration(key, value),
    }
}

func Time(key string, value time.Time) LoggerField {
    return LoggerField{
        field: zap.Time(key, value),
    }
}

func NewLogger() *Logger {
	zapLogger, _ := zap.NewProduction()
	return &Logger{
		logger: zapLogger,
	}
}

func (l *Logger) Debug(msg string, fields ...LoggerField) {
	var zapFields []zap.Field
	for _, field := range(fields) {
		zapFields = append(zapFields, field.field)
	}

	l.logger.Debug(msg, zapFields...)
}

func (l *Logger) Error(msg string, fields ...LoggerField) {
	var zapFields []zap.Field
	for _, field := range(fields) {
		zapFields = append(zapFields, field.field)
	}
	
	l.logger.Error(msg, zapFields...)
}

func (l *Logger) Fatal(msg string, fields ...LoggerField) {
	var zapFields []zap.Field
	for _, field := range(fields) {
		zapFields = append(zapFields, field.field)
	}
	
	l.logger.Fatal(msg, zapFields...)
}

func (l *Logger) Info(msg string, fields ...LoggerField) {
	var zapFields []zap.Field
	for _, field := range(fields) {
		zapFields = append(zapFields, field.field)
	}
	
	l.logger.Info(msg, zapFields...)
}

func (l *Logger) Panic(msg string, fields ...LoggerField) {
	var zapFields []zap.Field
	for _, field := range(fields) {
		zapFields = append(zapFields, field.field)
	}
	
	l.logger.Panic(msg, zapFields...)
}

func (l *Logger) Warn(msg string, fields ...LoggerField) {
	var zapFields []zap.Field
	for _, field := range(fields) {
		zapFields = append(zapFields, field.field)
	}
	
	l.logger.Warn(msg, zapFields...)
}

func (l *Logger) DPanic(msg string, fields ...LoggerField) {
	var zapFields []zap.Field
	for _, field := range(fields) {
		zapFields = append(zapFields, field.field)
	}
	
	l.logger.DPanic(msg, zapFields...)
}