package dish

import "time"

type Dish struct {
	Id          int64     `gorm:"column:id;primaryKey;type:bigint,autoIncrement"`
	Name        string    `gorm:"column:name;type:varchar(32)"`
	CategoryId  int64     `gorm:"column:category_id;type:bigint"`
	Price       float64   `gorm:"column:price;type:decimal(10,2)"`
	Image       string    `gorm:"column:image;type:varchar(255)"`
	Description string    `gorm:"column:description;type:varchar(255)"`
	Status      int       `gorm:"column:status;type:int"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime"`
	CreateUser  int64     `gorm:"column:create_user;type:bigint"`
	UpdateUser  int64     `gorm:"column:update_user;type:bigint"`
}

func (Dish) TableName() string {
	return "dish"
}
