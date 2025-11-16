package application

type SubmitDTO struct {
	AddressBookId         int64   `json:"addressBookId"`
	Amount                float64 `json:"amount"`
	DeliveryStatus        int     `json:"deliveryStatus"`
	EstimatedDeliveryTime string  `json:"estimatedDeliveryTime"`
	PackAmount            int     `json:"packAmount"`
	PayMethod             int     `json:"payMethod"`
	Remark                string  `json:"remark"`
	TablewareNumber       int     `json:"tablewareNumber"`
	TablewareStatus       int     `json:"tablewareStatus"`
}

type PaymentDTO struct {
	OrderNumber string `json:"orderNumber"`
	PayMethod   int    `json:"payMethod"`
}

type PageDTO struct {
	PageSize int `form:"pageSize"`
	Page     int `form:"page"`
	Status   int `form:"status"`
}
