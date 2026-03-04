package p

import "log/slog"

func _() {
	slog.Debug("")
	slog.Debug("good msg")
	slog.Debug(`good msg`)

	slog.Debug("\"")                     // want `message must not contain special symbols`
	slog.Debug("Bad msg")                // want `first letter must be in lower case`
	slog.Debug("русское сообщение")      // want `message must contain only english letters`
	slog.Debug("bad emoji message ✅✅✅✅") // want `message must not contain special symbols`
	slog.Debug("bad:special:message")    // want `message must not contain special symbols`
	slog.Debug("combination...✅")        // want `message must not contain special symbols`
	slog.Debug("bad few\nlines msg")     // want `message must not contain special symbols`
	slog.Debug(`"`)                      // want `message must not contain special symbols`
}
