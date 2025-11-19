package application

type SetMealDTO struct {
	CategoryId    int64         `json:"categoryId"`
	Description   string        `json:"description"`
	Id            int64         `json:"id"`
	Image         string        `json:"image"`
	Name          string        `json:"name"`
	Price         float64       `json:"price,string"`
	SetMealDishes []SetMealDish `json:"setmealDishes"`
	Status        int           `json:"status"`
}

type SetMealDish struct {
	Copies    int     `json:"copies"`
	DishId    int64   `json:"dishId"`
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	SetmealId int64   `json:"setmealId"`
}

type PageDTO struct {
	CategoryId int64  `form:"categoryId"`
	Name       string `form:"name"`
	Page       int    `form:"page"`
	PageSize   int    `form:"pageSize"`
	Status     string `form:"status"`
}
