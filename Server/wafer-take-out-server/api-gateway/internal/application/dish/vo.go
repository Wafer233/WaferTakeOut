package application

type PageVO struct {
	Total   int64    `json:"total"`
	Records []Record `json:"records"`
}

type Record struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	CategoryId   int64   `json:"categoryId"`
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
	Description  string  `json:"description"`
	Status       int     `json:"status"`
	UpdateTime   string  `json:"updateTime"`
	CategoryName string  `json:"categoryName"`
}

type DishVO struct {
	CategoryId   int64    `json:"categoryId"`
	CategoryName string   `json:"categoryName"`
	Description  string   `json:"description"`
	Flavors      []Flavor `json:"flavors"`
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`

	//这个地方不是string ！！ debug了几个小时
	Price      float64 `json:"price"`
	Status     int     `json:"status"`
	UpdateTime string  `json:"updateTime"`
}
