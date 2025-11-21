package logs

import "go.uber.org/zap"

func ReplaceGlobals(logger *zap.Logger) {
	zap.ReplaceGlobals(logger)
}
