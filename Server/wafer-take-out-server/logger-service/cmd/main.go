package main

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/logger-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/logger-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/logger-service/internal/infrastructure/kafka"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/logger-service/internal/infrastructure/persistence"
)

func main() {
	consumer := kafka.NewConsumer(
		[]string{"127.0.0.1:9092"},
		"wafer-take-out-log",
	)

	db, _ := database.NewMysqlDatabase()
	repo := persistence.NewDefaultLogMsgRepository(db)
	svc := application.NewLogService(repo, consumer)

	svc.Start()
}
