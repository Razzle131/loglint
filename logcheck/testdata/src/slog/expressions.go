package p

import "log/slog"

func _() {
	foo := "abc"
	slog.Debug("api1 " + "api2 " + "api3 " + "api4 ")
	slog.Debug("api1 " + foo + "api3 " + "api4 ")
	slog.Debug(foo + "abc")

	apiKey := "aboba"
	slog.Debug("api1" + apiKey)            // want `must not log sensitive data`
	slog.Debug("api1 " + apiKey + "api3 ") // want `must not log sensitive data`
	slog.Debug("апи ключ " + apiKey)       // want `must not log sensitive data` `message must contain only english letters`
	slog.Debug(apiKey + "foo")             // want `must not log sensitive data`
}
