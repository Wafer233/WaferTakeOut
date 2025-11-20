package addressbookApp

type AddressDTO struct {
	CityCode     string `json:"cityCode"`
	CityName     string `json:"cityName"`
	Consignee    string `json:"consignee"`
	Detail       string `json:"detail"`
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
	ID           int64  `json:"id"`
	IsDefault    int    `json:"isDefault"`
	Label        int    `json:"label"` //这个前端是int
	Phone        string `json:"phone"`
	ProvinceCode string `json:"provinceCode"`
	ProvinceName string `json:"provinceName"`
	Sex          string `json:"sex"`
	UserID       int64  `json:"userId"`
}

type DefaultIdDTO struct {
	Id int64 `json:"id"`
}
