package orderApp

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

type UserPageDTO struct {
	PageSize int `form:"pageSize"`
	Page     int `form:"page"`
	Status   int `form:"status"`
}

type AdminPageDTO struct {
	BeginTime string `form:"beginTime"`
	EndTime   string `form:"endTime"`
	Number    string `form:"number"`
	PageSize  int    `form:"pageSize"`
	Page      int    `form:"page"`
	Phone     string `form:"phone"`
	Status    int    `form:"status"`
}

type ConfirmDTO struct {
	Id int64 `json:"id"`
}

type RejectionDTO struct {
	RejectionReason string `json:"rejectionReason"`
	Id              int64  `json:"id"`
}

type CancelDTO struct {
	CancelReason string `json:"cancelReason"`
	Id           int64  `json:"id"`
}
