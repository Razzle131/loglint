package p

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func _() {
	apiKey := "abc"

	logger := zap.L()
	sugar := zap.L().Sugar()

	logger.Debug(apiKey)                   // want `must not log sensitive data`
	logger.Info(apiKey)                    // want `must not log sensitive data`
	logger.Warn(apiKey)                    // want `must not log sensitive data`
	logger.Error(apiKey)                   // want `must not log sensitive data`
	logger.Log(zapcore.DebugLevel, apiKey) // want `must not log sensitive data`

	sugar.Debug(apiKey)                   // want `must not log sensitive data`
	sugar.Info(apiKey)                    // want `must not log sensitive data`
	sugar.Warn(apiKey)                    // want `must not log sensitive data`
	sugar.Error(apiKey)                   // want `must not log sensitive data`
	sugar.Log(zapcore.DebugLevel, apiKey) // want `must not log sensitive data`

	sugar.Debugln(apiKey)                   // want `must not log sensitive data`
	sugar.Infoln(apiKey)                    // want `must not log sensitive data`
	sugar.Warnln(apiKey)                    // want `must not log sensitive data`
	sugar.Errorln(apiKey)                   // want `must not log sensitive data`
	sugar.Logln(zapcore.DebugLevel, apiKey) // want `must not log sensitive data`

	sugar.Debugw(apiKey)                   // want `must not log sensitive data`
	sugar.Infow(apiKey)                    // want `must not log sensitive data`
	sugar.Warnw(apiKey)                    // want `must not log sensitive data`
	sugar.Errorw(apiKey)                   // want `must not log sensitive data`
	sugar.Logw(zapcore.DebugLevel, apiKey) // want `must not log sensitive data`
}
