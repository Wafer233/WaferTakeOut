package category

import "time"

type Category struct {
	ID         int64     `gorm:"column:id;primary_key;type:bigint;not null"`
	Type       int       `gorm:"column:type;type:int"`
	Name       string    `gorm:"column:name;type:varchar(32);not null"`
	Sort       int       `gorm:"column:sort;type:int;not null"`
	Status     int       `gorm:"column:status;type:int"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime"`
	CreateUser int64     `gorm:"column:create_user;type:bigint"`
	UpdateUser int64     `gorm:"column:update_user;type:bigint"`
}

func (Category) TableName() string {
	return "category"
}
