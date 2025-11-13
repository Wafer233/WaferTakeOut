package application

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/flavor"
)

func (svc *DishService) UpdateDish(ctx context.Context,
	dto *DishDTO, curID int64) error {

	entity := &domain.Dish{
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

	if len(dto.Flavors) == 0 {
		return nil
	}

	flavorsDTO := dto.Flavors

	flavors := make([]*flavor.Flavor, len(flavorsDTO))
	for i, f := range flavorsDTO {
		flavors[i] = &flavor.Flavor{
			Id:     f.Id,
			DishId: entity.Id,
			Name:   f.Name,
			Value:  f.Value,
		}
	}
	err = svc.flavRepo.UpdatesByDishId(ctx, flavors)
	return nil
}
