package setmealApp

import (
	"context"
)

type GetSetMealVO struct {
	CategoryId    int64         `json:"categoryId"`
	CategoryName  string        `json:"categoryName"`
	Description   string        `json:"description"`
	Id            int64         `json:"id"`
	Image         string        `json:"image"`
	Name          string        `json:"name"`
	Price         float64       `json:"price,string"`
	SetMealDishes []SetMealDish `json:"setmealDishes"`
	Status        int           `json:"status"`
	UpdateTime    string        `json:"updateTime"`
}

func (svc *SetMealService) IdQuery(ctx context.Context, id int64) (GetSetMealVO, error) {

	set, err := svc.setRepo.GetById(ctx, id)
	if err != nil {
		return GetSetMealVO{}, err
	}
	dishes, err := svc.dishRepo.GetsBySetMealId(ctx, set.Id)
	if err != nil {
		return GetSetMealVO{}, err
	}

	dishVOs := make([]SetMealDish, len(dishes))

	for i, dish := range dishes {
		dishVOs[i] = SetMealDish{
			Copies:    dish.Copies,
			DishId:    dish.DishId,
			ID:        dish.Id,
			Name:      dish.Name,
			Price:     dish.Price,
			SetmealId: set.Id,
		}
	}

	category, err := svc.cateRepo.GetById(ctx, set.CategoryId)
	if err != nil {
		return GetSetMealVO{}, err
	}

	vo := GetSetMealVO{
		CategoryId:    set.CategoryId,
		CategoryName:  category.Name,
		Description:   set.Description,
		Id:            set.Id,
		Image:         set.Image,
		Name:          set.Name,
		Price:         set.Price,
		SetMealDishes: dishVOs,
		Status:        set.Status,
		UpdateTime:    set.UpdateTime.Format("2006-01-02 15:04"),
	}

	return vo, err

}
