package p

import (
	"go.uber.org/zap"
)

func _() {
	foo := "abc"
	zap.L().Debug("api1 " + "api2 " + "api3 " + "api4 ")
	zap.L().Debug("api1 " + foo + "api3 " + "api4 ")
	zap.L().Debug(foo + "abc")

	apiKey := "aboba"
	zap.L().Debug("api " + apiKey)            // want `must not log sensitive data`
	zap.L().Debug("api1 " + apiKey + "api3 ") // want `must not log sensitive data`
	zap.L().Debug("апи ключ " + apiKey)       // want `must not log sensitive data, message must contain only english letters`
	zap.L().Debug(apiKey + "foo")             // want `must not log sensitive data`
}
