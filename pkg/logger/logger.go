package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

type LoggerField struct {
	field zap.Field
}

func(f *LoggerField) GetLoggerField() zap.Field {
	return f.field
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