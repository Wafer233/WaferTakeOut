package main

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/initialize"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/infrastructure/kafka"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func main() {

	logger.Init()

	producer := kafka.NewProducer("127.0.0.1:9092")
	logger.InitKafkaLog(producer)

	defer producer.Close()

	r, err := initialize.Init()
	if err != nil {
		zap.L().Error("初始化失败: ", zap.Error(err))
		panic(err)
	}

	zap.L().Info("初始化成功")
	logger.K().Info("APP初始化成功")
	_ = r.Run(":8080")

}
