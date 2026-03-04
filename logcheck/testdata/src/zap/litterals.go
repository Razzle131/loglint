package p

import "go.uber.org/zap"

func _() {
	zap.L().Debug("")
	zap.L().Debug("good msg")
	zap.L().Debug(`good msg`)

	zap.L().Debug("\"")                     // want `message must not contain special symbols`
	zap.L().Debug("Bad msg")                // want `first letter must be in lower case`
	zap.L().Debug("русское сообщение")      // want `message must contain only english letters`
	zap.L().Debug("bad emoji message ✅✅✅✅") // want `message must not contain special symbols`
	zap.L().Debug("bad:special:message")    // want `message must not contain special symbols`
	zap.L().Debug("combination...✅")        // want `message must not contain special symbols`
	zap.L().Debug(`bad few 
lines msg`) // want `message must not contain special symbols`
	zap.L().Debug(`"`) // want `message must not contain special symbols`

}
