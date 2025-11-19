package application

type RecordVO struct {
	Amount     float64 `json:"amount"`
	CreateTime string  `json:"createTime"`
	DishFlavor string  `json:"dishFlavor"`
	DishId     int64   `json:"dishId"`
	Id         int64   `json:"id"`
	Image      string  `json:"image"`
	Name       string  `json:"name"`
	Number     int     `json:"number"`
	SetMealId  int64   `json:"setmealId"`
	UserID     int64   `json:"userId"`
}
