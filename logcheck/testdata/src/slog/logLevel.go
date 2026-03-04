package p

import (
	"context"
	"log/slog"
)

func _() {
	apiKey := "abc"

	slog.Log(context.Background(), slog.LevelDebug, apiKey)      // want `must not log sensitive data`
	slog.LogAttrs(context.Background(), slog.LevelDebug, apiKey) // want `must not log sensitive data`

	slog.Debug(apiKey) // want `must not log sensitive data`
	slog.Info(apiKey)  // want `must not log sensitive data`
	slog.Warn(apiKey)  // want `must not log sensitive data`
	slog.Error(apiKey) // want `must not log sensitive data`

	slog.DebugContext(context.Background(), apiKey) // want `must not log sensitive data`
	slog.InfoContext(context.Background(), apiKey)  // want `must not log sensitive data`
	slog.WarnContext(context.Background(), apiKey)  // want `must not log sensitive data`
	slog.ErrorContext(context.Background(), apiKey) // want `must not log sensitive data`

	logger := &slog.Logger{}
	logger.Log(context.Background(), slog.LevelDebug, apiKey)      // want `must not log sensitive data`
	logger.LogAttrs(context.Background(), slog.LevelDebug, apiKey) // want `must not log sensitive data`

	logger.Debug(apiKey) // want `must not log sensitive data`
	logger.Info(apiKey)  // want `must not log sensitive data`
	logger.Warn(apiKey)  // want `must not log sensitive data`
	logger.Error(apiKey) // want `must not log sensitive data`

	logger.DebugContext(context.Background(), apiKey) // want `must not log sensitive data`
	logger.InfoContext(context.Background(), apiKey)  // want `must not log sensitive data`
	logger.WarnContext(context.Background(), apiKey)  // want `must not log sensitive data`
	logger.ErrorContext(context.Background(), apiKey) // want `must not log sensitive data`
}
