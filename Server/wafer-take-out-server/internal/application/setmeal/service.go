package setmealApp

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal_dish"
)

type SetMealService struct {
	setRepo  setmeal.SetMealRepository
	dishRepo setmeal_dish.SetMealDishRepository
	cateRepo category.CategoryRepository
}

func NewSetMealService(setRepo setmeal.SetMealRepository,
	dishRepo setmeal_dish.SetMealDishRepository,
	cateRepo category.CategoryRepository) *SetMealService {
	return &SetMealService{
		setRepo:  setRepo,
		dishRepo: dishRepo,
		cateRepo: cateRepo,
	}
}
