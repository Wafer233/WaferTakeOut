package application

type SubmitVO struct {
	Id          int64   `json:"id"`
	OrderAmount float64 `json:"orderAmount"`
	OrderNumber string  `json:"orderNumber"`
	OrderTime   string  `json:"orderTime"`
}

type PaymentVO struct {
	EstimatedDeliveryTime string `json:"estimatedDeliveryTime"`
}
