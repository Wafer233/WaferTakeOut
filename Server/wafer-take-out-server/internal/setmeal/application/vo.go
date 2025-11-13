package application

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

type PageVO struct {
	Total   int64    `json:"total"`
	Records []Record `json:"records"`
}

type Record struct {
	Id           int64   `json:"id"`
	CategoryId   int64   `json:"categoryId"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
	Description  string  `json:"description"`
	Image        string  `json:"image"`
	UpdateTime   string  `json:"updateTime"`
	CategoryName string  `json:"categoryName"`
}
