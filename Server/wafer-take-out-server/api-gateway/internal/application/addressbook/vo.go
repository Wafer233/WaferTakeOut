package addressbookApp

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
