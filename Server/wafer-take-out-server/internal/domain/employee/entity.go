package employee

import "time"

type Employee struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string    `gorm:"column:name;type:varchar(32);not null"`
	Username   string    `gorm:"column:username;type:varchar(32);not null"`
	Password   string    `gorm:"column:password;type:varchar(64);not null"`
	Phone      string    `gorm:"column:phone;type:varchar(11);not null"`
	Sex        string    `gorm:"column:sex;type:varchar(2);not null"`
	IDNumber   string    `gorm:"column:id_number;type:varchar(18);not null"`
	Status     int       `gorm:"column:status;type:int;not null;default:1"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`
	CreateUser int64     `gorm:"column:create_user;type:bigint"`
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`
}

func (Employee) TableName() string {
	return "employee"
}
