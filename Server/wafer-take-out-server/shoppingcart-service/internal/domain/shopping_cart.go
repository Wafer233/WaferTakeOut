package domain

import "time"

type ShoppingCart struct {
	Id         int64     `gorm:"column:id;primaryKey;type:bigint,autoIncrement"`
	Name       string    `gorm:"column:name;type:varchar(32);"`
	Image      string    `gorm:"column:image;type:varchar(255);"`
	UserId     int64     `gorm:"column:user_id;type:bigint;"`
	DishId     int64     `gorm:"column:dish_id;type:bigint;"`
	SetmealId  int64     `gorm:"column:setmeal_id;type:bigint;"`
	DishFlavor string    `gorm:"column:dish_flavor;type:varchar(50);"`
	Number     int       `gorm:"column:number;type:int;"`
	Amount     float64   `gorm:"column:amount;type:decimal(10,2);"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;"`
}

func (ShoppingCart) TableName() string {
	return "shopping_cart"
}
