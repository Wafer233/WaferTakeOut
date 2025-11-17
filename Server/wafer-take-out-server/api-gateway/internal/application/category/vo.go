package categoryApp

type PageVO struct {
	Total   int64    `json:"total"`
	Records []Record `json:"records"`
}
type Record struct {
	ID         int64  `json:"id"`
	Type       int    `json:"type"`
	Name       string `json:"name"`
	Sort       int    `json:"sort"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	CreateUser int64  `json:"createUser"`
	UpdateUser int64  `json:"updateUser"`
}
