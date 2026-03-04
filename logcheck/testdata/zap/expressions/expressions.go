package testdata

import (
	"go.uber.org/zap"
)

func _() {

	// good
	foo := "abc"
	zap.L().Debug("api1 " + "api2 " + "api3 " + "api4 ")
	zap.L().Debug("api1 " + foo + "api3 " + "api4 ")
	zap.L().Debug(foo + "abc")
	// ------------
	// bad
	apiKey := "aboba"
	zap.L().Debug("api " + apiKey)            // sensetive name - apiKey
	zap.L().Debug("api1 " + apiKey + "api3 ") // sensetive name - apiKey
	zap.L().Debug("апи ключ " + apiKey)       // sensetive name - apiKey, non-english language
	zap.L().Debug(apiKey + "foo")             // sensetive name - apiKey
}
