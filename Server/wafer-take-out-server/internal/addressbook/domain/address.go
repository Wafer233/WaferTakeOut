package domain

type AddressBook struct {
	Id           int64  `gorm:"column:id;primaryKey;type:bigint,autoIncrement"`
	UserId       int64  `gorm:"column:user_id;primaryKey;type:bigint"`
	Consignee    string `gorm:"column:consignee;type:varchar(50)"`
	Sex          string `gorm:"column:sex;varchar(2)"`
	Phone        string `gorm:"column:phone;varchar(11)"`
	ProvinceCode string `gorm:"column:province_code;varchar(12)"`
	ProvinceName string `gorm:"column:province_name;varchar(32)"`
	CityCode     string `gorm:"column:city_code;varchar(12)"`
	CityName     string `gorm:"column:city_name;varchar(32)"`
	DistrictCode string `gorm:"column:district_code;varchar(12)"`
	DistrictName string `gorm:"column:district_name;varchar(32)"`
	Detail       string `gorm:"column:detail;varchar(200)"`
	Label        string `gorm:"column:label;varchar(100)"`
	IsDefault    int    `gorm:"column:is_default;type:tinyint(1)"`
}

func (AddressBook) TableName() string {
	return "address_book"
}
