package domain

import "time"

type User struct {
	Id         int64     `gorm:"column:id;primaryKey;type:bigint,autoIncrement"`
	OpenId     string    `gorm:"column:openid;type:varchar(32)"`
	Name       string    `gorm:"column:name;type:varchar(11)"`
	Phone      string    `gorm:"column:phone;type:varchar(11)"`
	Sex        string    `gorm:"column:sex;type:varchar(2)"`
	IdNumber   string    `gorm:"column:id_number;type:varchar(18)"`
	Avatar     string    `gorm:"column:avatar;type:varchar(500)"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime"`
}

func (User) TableName() string {
	return "user"
}
