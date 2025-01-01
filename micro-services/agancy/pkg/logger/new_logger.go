package logger

import (
	"context"
	"fmt"
	"gholi-fly-agancy/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerService struct {
	logger *zap.Logger
}

// NewLoggerService initializes the logger service for a single service.
func NewLoggerService(cfg config.LoggerConfig) (*LoggerService, error) {
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.LogFilePath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})

	atomicLevel := zap.NewAtomicLevel()
	if err := atomicLevel.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, fmt.Errorf("invalid log level: %v", err)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		fileWriter,
		atomicLevel,
	)

	logger := zap.New(core)
	return &LoggerService{logger: logger}, nil
}

// LogInfo logs informational messages.
func (l *LoggerService) LogInfo(ctx context.Context, message string, fields ...zap.Field) {
	l.logger.Info(message, fields...)
}

// LogError logs error messages.
func (l *LoggerService) LogError(ctx context.Context, message string, fields ...zap.Field) {
	l.logger.Error(message, fields...)
}

// Sync flushes any buffered log entries.
func (l *LoggerService) Sync() {
	_ = l.logger.Sync()
}

// AttachLoggerToContext adds logger to the context for Fiber.
func (l *LoggerService) AttachLoggerToContext(ctx context.Context, fiberCtx *fiber.Ctx, reqID string) context.Context {
	fields := []zap.Field{
		zap.String("method", fiberCtx.Method()),
		zap.String("url", fiberCtx.OriginalURL()),
		zap.String("ip", fiberCtx.IP()),
		zap.String("user_agent", fiberCtx.Get("User-Agent")),
		zap.String("X-Request-ID", reqID),
	}
	return context.WithValue(ctx, "logger", l.logger.With(fields...))
}
