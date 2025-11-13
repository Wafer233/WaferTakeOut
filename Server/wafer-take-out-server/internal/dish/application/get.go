package application

import (
	"context"
)

type DishVO struct {
	CategoryId   int64    `json:"categoryId"`
	CategoryName string   `json:"categoryName"`
	Description  string   `json:"description"`
	Flavors      []Flavor `json:"flavors"`
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Price        float64  `json:"price,string"`
	Status       int      `json:"status"`
	UpdateTime   string   `json:"updateTime"`
}

func (svc *DishService) GetDishId(ctx context.Context, id int64) (DishVO, error) {

	dish, err := svc.dishRepo.GetById(ctx, id)
	if err != nil {
		return DishVO{}, err
	}
	category, err := svc.cateRepo.GetById(ctx, dish.CategoryId)
	if err != nil {
		return DishVO{}, err
	}

	flavors, err := svc.flavRepo.GetsByDishId(ctx, dish.Id)
	if err != nil {
		return DishVO{}, err
	}

	flavorVOs := make([]Flavor, len(flavors))
	for index, flavor := range flavors {
		flavorVOs[index].DishId = flavor.DishId
		flavorVOs[index].Value = flavor.Value
		flavorVOs[index].Name = flavor.Name
		flavorVOs[index].Id = flavor.Id
	}

	vo := DishVO{
		CategoryId:   dish.CategoryId,
		CategoryName: category.Name,
		Description:  dish.Description,
		Flavors:      flavorVOs,
		ID:           dish.Id,
		Image:        dish.Image,
		Name:         dish.Name,
		Price:        dish.Price,
		Status:       dish.Status,
		UpdateTime:   dish.UpdateTime.Format("2006-01-02 15:04"),
	}
	return vo, nil
}
