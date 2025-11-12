package setmeal_dish

type SetMealDish struct {
	Id        int64   `gorm:"primaryKey;column:id;type:bigint,autoIncrement"`
	SetMealId int64   `gorm:"column:setmeal_id;type:bigint"`
	DishId    int64   `gorm:"column:dish_id;type:bigint"`
	Name      string  `gorm:"column:name;type:varchar(32);"`
	Price     float64 `gorm:"column:price;type:decimal(10,2);"`
	Copies    int     `gorm:"column:copies;type:int"`
}

func (SetMealDish) TableName() string {
	return "setmeal_dish"
}
