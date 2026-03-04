package testdata

import "go.uber.org/zap"

func _() {
	// good
	foo := "abc"
	zap.L().Debug(foo)
	// ------------
	// bad
	token := "abc"
	zap.L().Debug(token) // sensetive name - token
}
