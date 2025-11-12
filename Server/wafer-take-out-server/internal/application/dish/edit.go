package dishApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/dish"
)

func (svc *DishService) UpdateDish(ctx context.Context,
	dto *DishDTO, curID int64) error {

	entity := &dish.Dish{
		Id:          dto.ID,
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       dto.Price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      dto.Status,
		UpdateTime:  time.Now(),
		UpdateUser:  curID,
	}
	err := svc.dishRepo.UpdateById(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}
