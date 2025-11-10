package employeeApp

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/times"
)

type PageVO struct {
	Total     int64      `json:"total"`
	Employees []Employee `json:"records"`
}

type Employee struct {
	Id         int64          `json:"id"`
	Name       string         `json:"name"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Phone      string         `json:"phone"`
	Sex        string         `json:"sex"`
	IDNumber   string         `json:"idNumber"`
	Status     int            `json:"status"`
	CreateTime times.JSONTime `json:"createTime"`
	UpdateTime times.JSONTime `json:"updateTime"`
	CreateUser int64          `json:"createUser"`
	UpdateUser int64          `json:"updateUser"`
}
