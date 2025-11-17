package application

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

type PageDTO struct {
	CategoryId int64  `form:"categoryId"`
	Name       string `form:"name"`
	Page       int    `form:"page"`
	PageSize   int    `form:"pageSize"`
	Status     string `form:"status"`
}
