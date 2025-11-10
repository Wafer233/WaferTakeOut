package employeeApp

type AddEmployeeDTO struct {
	ID       int64  `json:"id"`
	IDNumber string `json:"idNumber"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	UserName string `json:"username"`
}
