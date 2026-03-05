package p

import (
	"log/slog"
)

func _() {
	apiKey := "abc"

	slog.Debug(apiKey) // want `must not log sensitive data`
	slog.Info(apiKey)  // want `must not log sensitive data`
	slog.Warn(apiKey)  // want `must not log sensitive data`
	slog.Error(apiKey) // want `must not log sensitive data`
}
