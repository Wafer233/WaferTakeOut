package setmealApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal_dish"
)

// 这里的price又是string
type AddSetMealDTO struct {
	CategoryId    int64         `json:"categoryId"`
	Description   string        `json:"description"`
	Id            int64         `json:"id"`
	Image         string        `json:"image"`
	Name          string        `json:"name"`
	Price         float64       `json:"price,string"`
	SetMealDishes []SetMealDish `json:"setmealDishes"`
	Status        int           `json:"status"`
}

type SetMealDish struct {
	Copies    int     `json:"copies"`
	DishId    int64   `json:"dishId"`
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	SetmealId int64   `json:"setmealId"`
}

func (svc *SetMealService) AddSetMeal(ctx context.Context, dto *AddSetMealDTO, curId int64) error {

	set := &setmeal.SetMeal{
		Id:          dto.Id,
		CategoryId:  dto.CategoryId,
		Name:        dto.Name,
		Price:       dto.Price,
		Status:      dto.Status,
		Description: dto.Description,
		Image:       dto.Image,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		CreateUser:  curId,
		UpdateUser:  curId,
	}

	mealDishes := make([]*setmeal_dish.SetMealDish, len(dto.SetMealDishes))

	err := svc.setRepo.Insert(ctx, set)
	if err != nil {
		return err
	}

	for index, value := range dto.SetMealDishes {
		mealDishes[index] = &setmeal_dish.SetMealDish{
			Id:        value.ID,
			SetMealId: set.Id,
			DishId:    value.DishId,
			Name:      value.Name,
			Price:     value.Price,
			Copies:    value.Copies,
		}
	}

	err = svc.dishRepo.Inserts(ctx, mealDishes)
	if err != nil {
		return err
	}
	return nil
}
