package application

import (
	"fmt"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/logger-service/internal/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/logger-service/internal/infrastructure/kafka"
	"github.com/jinzhu/copier"
)

type LogService struct {
	repo     domain.LogMsgRepository
	consumer *kafka.Consumer
}

func NewLogService(
	repo domain.LogMsgRepository,
	consumer *kafka.Consumer,
) *LogService {
	return &LogService{
		repo:     repo,
		consumer: consumer,
	}
}

func (svc *LogService) Start() {
	for {
		// 1. 从 Kafka 拉取
		rec, err := svc.consumer.Fetch()
		if err != nil {
			fmt.Println("拉取失败:", err)
			continue
		}

		// 2. 转为领域对象
		var entity domain.LogMsg
		_ = copier.Copy(&entity, rec)
		layout := "2006-01-02 15:04:05"
		tme, _ := time.Parse(layout, rec.Time)
		entity.Time = tme
		// 3. 存储
		if err := svc.repo.Save(&entity); err != nil {
			fmt.Println("存储失败:", err)
		}
	}
}
