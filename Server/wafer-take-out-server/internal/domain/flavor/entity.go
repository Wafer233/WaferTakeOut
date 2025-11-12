package flavor

type Flavor struct {
	Id     int64  `gorm:"column:id;primaryKey;type:bigint,autoIncrement"`
	DishId int64  `gorm:"column:dish_id;type:bigint"`
	Name   string `gorm:"column:name;type:varchar(32)"`
	Value  string `gorm:"column:value;type:varchar(255)"`
}

func (Flavor) TableName() string {
	return "dish_flavor"
}
