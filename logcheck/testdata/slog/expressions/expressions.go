package testdata

import "log/slog"

func _() {
	// good
	foo := "abc"
	slog.Debug("api1 " + "api2 " + "api3 " + "api4 ")
	slog.Debug("api1 " + foo + "api3 " + "api4 ")
	slog.Debug(foo + "abc")
	// ------------
	// bad
	apiKey := "aboba"
	slog.Debug("api " + apiKey)            // sensetive name - apiKey
	slog.Debug("api1 " + apiKey + "api3 ") // sensetive name - apiKey
	slog.Debug("апи ключ " + apiKey)       // sensetive name - apiKey, non-english language
	slog.Debug(apiKey + "foo")             // sensetive name - apiKey
}
