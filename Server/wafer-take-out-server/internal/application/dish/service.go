package dishApp

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/flavor"
)

type DishService struct {
	dishRepo dish.DishRepository
	cateRepo category.CategoryRepository
	flavRepo flavor.FlavorRepository
}

func NewDishService(dishRepo dish.DishRepository,
	cateRepo category.CategoryRepository,
	flavRepo flavor.FlavorRepository) *DishService {
	return &DishService{
		dishRepo: dishRepo,
		cateRepo: cateRepo,
		flavRepo: flavRepo,
	}
}
