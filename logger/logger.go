package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log ..
var log *zap.Logger

func init() {
	var err error
	c := zap.NewProductionConfig()
	ec := zap.NewProductionEncoderConfig()
	ec.TimeKey = "timestamp"
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	ec.StacktraceKey = ""
	c.EncoderConfig = ec
	log, err = c.Build(zap.AddCallerSkip(1))
	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// Info ..
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug ..
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error ..
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
