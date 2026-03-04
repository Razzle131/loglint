package testdata

import (
	"context"
	"log/slog"
)

func _() {
	apiKey := "abc"

	slog.Log(context.Background(), slog.LevelDebug, apiKey)
	slog.LogAttrs(context.Background(), slog.LevelDebug, apiKey)

	slog.Debug(apiKey)
	slog.Info(apiKey)
	slog.Warn(apiKey)
	slog.Error(apiKey)

	slog.DebugContext(context.Background(), apiKey)
	slog.InfoContext(context.Background(), apiKey)
	slog.WarnContext(context.Background(), apiKey)
	slog.ErrorContext(context.Background(), apiKey)

	logger := &slog.Logger{}
	logger.Log(context.Background(), slog.LevelDebug, apiKey)
	logger.LogAttrs(context.Background(), slog.LevelDebug, apiKey)

	logger.Debug(apiKey)
	logger.Info(apiKey)
	logger.Warn(apiKey)
	logger.Error(apiKey)

	logger.DebugContext(context.Background(), apiKey)
	logger.InfoContext(context.Background(), apiKey)
	logger.WarnContext(context.Background(), apiKey)
	logger.ErrorContext(context.Background(), apiKey)
}
