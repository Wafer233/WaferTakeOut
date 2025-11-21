package logs

import (
	"testing"

	"go.uber.org/zap"
)

func TestInit(t *testing.T) {
	Init()

	zap.L().Info("全局日志测试")
}
