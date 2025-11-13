package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/domain"
)

func (svc *SetMealService) StatusFlip(ctx context.Context, id int64,
	curId int64, status int) error {

	entity := &domain.SetMeal{
		Id:         id,
		Status:     status,
		UpdateUser: curId,
		UpdateTime: time.Now(),
	}
	err := svc.setRepo.UpdateStatusById(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}
