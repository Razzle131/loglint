package testdata

import "log/slog"

func _() {
	// good
	foo := "abc"
	slog.Debug(foo)
	// ------------
	// bad
	token := "abc"
	slog.Debug(token) // sensetive name - token
}
