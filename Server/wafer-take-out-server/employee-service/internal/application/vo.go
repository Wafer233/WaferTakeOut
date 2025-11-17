package application

type LoginVO struct {
	ID       int64  `json:"id"`       // 主键值
	UserName string `json:"userName"` // 用户名
	Name     string `json:"name"`     // 姓名
	Token    string `json:"token"`    // JWT令牌
}

type PageVO struct {
	Total     int64      `json:"total"`
	Employees []Employee `json:"records"`
}

type Employee struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Sex        string `json:"sex"`
	IDNumber   string `json:"idNumber"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	CreateUser int64  `json:"createUser"`
	UpdateUser int64  `json:"updateUser"`
}
