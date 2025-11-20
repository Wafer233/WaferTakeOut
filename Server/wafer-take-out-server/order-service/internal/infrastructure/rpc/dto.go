package rpc

type CartDTO struct {
	DishFlavor string `json:"dishFlavor"`
	DishId     int64  `json:"dishId"`
	SetMealId  int64  `json:"setmealId"`
}
