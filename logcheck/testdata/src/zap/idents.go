package p

import "go.uber.org/zap"

func _() {
	foo := "abc"
	zap.L().Debug(foo)

	token := "abc"
	zap.L().Debug(token) // want `must not log sensitive data`
}
