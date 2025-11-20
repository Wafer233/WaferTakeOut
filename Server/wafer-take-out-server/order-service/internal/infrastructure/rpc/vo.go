package rpc

type AddressBookVO struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"userId"`
	Consignee    string `json:"consignee"`
	Sex          string `json:"sex"`
	Phone        string `json:"phone"`
	ProvinceCode string `json:"provinceCode"`
	ProvinceName string `json:"provinceName"`
	CityCode     string `json:"cityCode"`
	CityName     string `json:"cityName"`
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
	Detail       string `json:"detail"`
	Label        string `json:"label"`
	IsDefault    int    `json:"isDefault"`
}

type CartVO struct {
	Amount     float64 `json:"amount"`
	CreateTime string  `json:"createTime"`
	DishFlavor string  `json:"dishFlavor"`
	DishId     int64   `json:"dishId"`
	Id         int64   `json:"id"`
	Image      string  `json:"image"`
	Name       string  `json:"name"`
	Number     int     `json:"number"`
	SetMealId  int64   `json:"setmealId"`
	UserID     int64   `json:"userId"`
}
