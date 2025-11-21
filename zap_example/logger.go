package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() *zap.Logger {

	core := NewTeeCore()

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	ReplaceGlobals(logger)

	return logger
}
