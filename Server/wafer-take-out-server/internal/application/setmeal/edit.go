package setmealApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal_dish"
)

func (svc *SetMealService) Edit(ctx context.Context,
	dto *AddSetMealDTO, curId int64) error {

	set := &setmeal.SetMeal{
		Id:          dto.Id,
		CategoryId:  dto.CategoryId,
		Name:        dto.Name,
		Price:       dto.Price,
		Description: dto.Description,
		Image:       dto.Image,
		UpdateTime:  time.Now(),
		UpdateUser:  curId,
	}
	err := svc.setRepo.UpdateById(ctx, set)
	if err != nil {
		return err
	}

	if len(dto.SetMealDishes) == 0 {
		return nil
	}

	sets := make([]*setmeal_dish.SetMealDish, len(dto.SetMealDishes))
	for i, dish := range dto.SetMealDishes {
		sets[i] = &setmeal_dish.SetMealDish{
			Id:        dish.ID,
			SetMealId: set.Id,
			DishId:    dish.DishId,
			Name:      dish.Name,
			Price:     dish.Price,
			Copies:    dish.Copies,
		}
	}
	err = svc.dishRepo.UpdatesBySetMealId(ctx, sets)
	if err != nil {
		return err
	}
	return nil
}
