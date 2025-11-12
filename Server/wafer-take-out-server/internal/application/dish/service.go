package dishApp

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/dish"
)

type DishService struct {
	dishRepo dish.DishRepository
	cateRepo category.CategoryRepository
}

func NewDishService(dishRepo dish.DishRepository, cateRepo category.CategoryRepository) *DishService {
	return &DishService{
		dishRepo: dishRepo,
		cateRepo: cateRepo,
	}
}
