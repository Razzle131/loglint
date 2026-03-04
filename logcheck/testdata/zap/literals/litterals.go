package testdata

import "go.uber.org/zap"

func _() {
	// good
	zap.L().Debug("")
	zap.L().Debug("good msg")
	zap.L().Debug(`good msg`)
	// ------------
	// bad
	zap.L().Debug("\"")                     // special symbol
	zap.L().Debug("Bad msg")                // first letter is capital
	zap.L().Debug("русское сообщение")      // non-english
	zap.L().Debug("bad emoji message ✅✅✅✅") // special symbol (emoji)
	zap.L().Debug("bad:special:message")    // special symbol
	zap.L().Debug("combination...✅")        // special symbol
	zap.L().Debug(`bad few 
lines msg`) // special symbol
	zap.L().Debug(`"`) // special symbol

}
