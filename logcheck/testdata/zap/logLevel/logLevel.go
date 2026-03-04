package testdata

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func _() {
	apiKey := "abc"

	logger := zap.L()
	sugar := zap.L().Sugar()

	logger.Debug(apiKey)
	logger.Info(apiKey)
	logger.Warn(apiKey)
	logger.Error(apiKey)
	logger.Log(zapcore.DebugLevel, apiKey)

	sugar.Debug(apiKey)
	sugar.Info(apiKey)
	sugar.Warn(apiKey)
	sugar.Error(apiKey)
	sugar.Log(zapcore.DebugLevel, apiKey)

	sugar.Debugln(apiKey)
	sugar.Infoln(apiKey)
	sugar.Warnln(apiKey)
	sugar.Errorln(apiKey)
	sugar.Logln(zapcore.DebugLevel, apiKey)

	sugar.Debugw(apiKey)
	sugar.Infow(apiKey)
	sugar.Warnw(apiKey)
	sugar.Errorw(apiKey)
	sugar.Logw(zapcore.DebugLevel, apiKey)
}
