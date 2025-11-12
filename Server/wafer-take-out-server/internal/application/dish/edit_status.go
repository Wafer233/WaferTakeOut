package dishApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/dish"
)

func (svc *DishService) StatusFlip(ctx context.Context, id int64, status int, curId int64) error {

	entity := &dish.Dish{
		Id:         id,
		Status:     status,
		UpdateTime: time.Now(),
		UpdateUser: curId,
	}
	err := svc.dishRepo.UpdateStatusById(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
