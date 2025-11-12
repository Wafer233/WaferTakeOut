package dishApp

import (
	"context"
	"time"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/domain/flavor"
)

type DishDTO struct {
	CategoryId  int64    `json:"categoryId"`
	Description string   `json:"description"`
	Flavors     []Flavor `json:"flavors"`
	ID          int64    `json:"id"`
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	Price       float64  `json:"price,string"`
	Status      int      `json:"status"`
}

type Flavor struct {
	DishId int64  `json:"dishId"`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

func (svc *DishService) Insert(ctx context.Context, dto *DishDTO, curId int64) error {

	entity := &dish.Dish{
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       dto.Price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      dto.Status,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		CreateUser:  curId,
		UpdateUser:  curId,
	}
	err := svc.dishRepo.Insert(ctx, entity)
	if err != nil {
		return err
	}

	if len(dto.Flavors) != 0 {
		entities := make([]*flavor.Flavor, len(dto.Flavors))
		for index, value := range entities {
			value = &flavor.Flavor{}
			value.DishId = entity.Id
			value.Name = dto.Flavors[index].Name
			value.Value = dto.Flavors[index].Value

			entities[index] = value
		}

		err = svc.flavRepo.Inserts(ctx, entities)
		if err != nil {
			return err
		}
	}
	return nil
}
