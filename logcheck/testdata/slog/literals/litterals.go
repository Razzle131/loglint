package testdata

import "log/slog"

func _() {
	// good
	slog.Debug("")
	slog.Debug("good msg")
	slog.Debug(`good msg`)
	// ------------
	// bad
	slog.Debug("\"")                     // special symbol
	slog.Debug("Bad msg")                // first letter is capital
	slog.Debug("русское сообщение")      // non-english
	slog.Debug("bad emoji message ✅✅✅✅") // special symbol (emoji)
	slog.Debug("bad:special:message")    // special symbol
	slog.Debug("combination...✅")        // special symbol
	slog.Debug(`bad few 
lines msg`) // special symbol
	slog.Debug(`"`) // special symbol

}
