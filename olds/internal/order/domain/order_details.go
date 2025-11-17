package domain

type OrderDetail struct {
	Id         int64   `gorm:"column:id;primaryKey;Type:bigint;autoIncrement"`
	Name       string  `gorm:"column:name;type:varchar(32)"`
	Image      string  `gorm:"column:image;type:varchar(255)"`
	OrderId    int64   `gorm:"column:order_id;Type:bigint"`
	DishId     int64   `gorm:"column:dish_id;Type:bigint"`
	SetMealId  int64   `gorm:"column:setmeal_id;Type:bigint"`
	DishFlavor string  `gorm:"column:dish_flavor;Type:varchar(50)"`
	Number     int     `gorm:"column:number;Type:int"`
	Amount     float64 `gorm:"column:amount;Type:decimal(10,2)"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}
