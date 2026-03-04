package p

import "log/slog"

func _() {
	foo := 123
	slog.Debug("abc", "foo", foo)
	slog.Debug("abc", "foo", 12345)
	slog.Debug("abc", "foo", "foo")

	password := 123
	slog.Debug("abc", "foo", password) // want `must not log sensitive data`
	slog.Debug("abc", "Foo", foo)      // want `first letter must be in lower case`
	slog.Debug("Abc", "foo", password) // want `first letter must be in lower case` `must not log sensitive data`
	slog.Debug("abc", password)        // want `must not log sensitive data`
}
