package p

import "log/slog"

func _() {
	foo := "abc"
	slog.Debug(foo)

	token := "abc"
	slog.Debug(token) // want `must not log sensitive data`
}
