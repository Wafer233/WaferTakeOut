package logger

import (
	"os"

	"go.uber.org/zap/zapcore"
)

// 返回：tee 核心 (console + file)
func NewTeeCore() zapcore.Core {
	consoleEnc := NewConsoleEncoder()
	fileEnc := NewFileEncoder()

	// console 输出
	consoleCore := zapcore.NewCore(
		consoleEnc,
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	// file 输出
	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileCore := zapcore.NewCore(
		fileEnc,
		zapcore.AddSync(file),
		zapcore.InfoLevel,
	)

	// tee（双写）
	return zapcore.NewTee(consoleCore, fileCore)
}
