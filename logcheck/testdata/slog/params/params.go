package testdata

import "log/slog"

func _() {
	// with params
	// good
	foo := 123
	slog.Debug("abc", "foo", foo)
	slog.Debug("abc", "foo", 12345)
	slog.Debug("abc", "foo", "foo")
	// ------------
	// bad
	password := 123
	slog.Debug("abc", "foo", password) // sensetive name - password
	slog.Debug("abc", "Foo", foo)      // arg name starts with capital letter
	slog.Debug("Abc", "foo", password) // msg starts with capital letter
	slog.Debug("abc", password)        // sensetive name - password
}
