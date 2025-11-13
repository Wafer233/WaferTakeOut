package application

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/domain"
	domain2 "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/flavor"
)

type DishService struct {
	dishRepo domain2.DishRepository
	cateRepo domain.CategoryRepository
	flavRepo flavor.FlavorRepository
}

func NewDishService(dishRepo domain2.DishRepository,
	cateRepo domain.CategoryRepository,
	flavRepo flavor.FlavorRepository) *DishService {
	return &DishService{
		dishRepo: dishRepo,
		cateRepo: cateRepo,
		flavRepo: flavRepo,
	}
}
