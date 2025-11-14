package application

type LoginDTO struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type PageDTO struct {
	Name     string `form:"name"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

type StatusFlipsDTO struct {
	ID int64 `form:"id"`
}

type AddEmployeeDTO struct {
	ID       int64  `json:"id"`
	IDNumber string `json:"idNumber"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	UserName string `json:"username"`
}

type PasswordDTO struct {
	ID          int64  `json:"empId"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}
