package application

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/setmeal_dish"
	domain2 "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/domain"
)

type SetMealService struct {
	setRepo  domain2.SetMealRepository
	dishRepo setmeal_dish.SetMealDishRepository
	cateRepo domain.CategoryRepository
}

func NewSetMealService(setRepo domain2.SetMealRepository,
	dishRepo setmeal_dish.SetMealDishRepository,
	cateRepo domain.CategoryRepository) *SetMealService {
	return &SetMealService{
		setRepo:  setRepo,
		dishRepo: dishRepo,
		cateRepo: cateRepo,
	}
}
