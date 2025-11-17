package application

// 这里我都要骂人了，json绑定一直失败，前端文档说传进来的是int实际上是string！！
type AddCategoryDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort,string"`
	Type int    `json:"type,string"`
}

// 这里前端有个bug，就是如果我不点击对应的栏，从新输入的话，他前端返回的sort就是int类型
// 如果点击了，就是string类型，导致我反序列化失败。这里我就用string。
type EditCategoryDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort,string"`
	//Type int    `json:"type,string"`
}

type PageDTO struct {
	Name     string `form:"name"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Type     int    `form:"type"`
}
