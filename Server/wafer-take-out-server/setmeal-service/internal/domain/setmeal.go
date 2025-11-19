package domain

import "time"

type SetMeal struct {
	Id          int64     `gorm:"primaryKey;column:id;type:bigint,autoIncrement"`
	CategoryId  int64     `gorm:"column:category_id;type:bigint"`
	Name        string    `gorm:"column:name;type:varchar(32)"`
	Price       float64   `gorm:"column:price;type:decimal(10,2)"`
	Status      int       `gorm:"column:status;type:int"`
	Description string    `gorm:"column:description;type:varchar(255)"`
	Image       string    `gorm:"column:image;type:varchar(255)"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime"`
	CreateUser  int64     `gorm:"column:create_user;type:bigint"`
	UpdateUser  int64     `gorm:"column:update_user;type:bigint"`
}

func (SetMeal) TableName() string {
	return "setmeal"
}
