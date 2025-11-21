package main

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/initialize"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/logs"
	"go.uber.org/zap"
)

func main() {

	logs.Init()

	r, err := initialize.Init()
	if err != nil {
		zap.L().Error("初始化失败: ", zap.Error(err))
		panic(err)
	}

	_ = r.Run(":8080")

}
