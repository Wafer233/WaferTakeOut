package orderApp

type SubmitVO struct {
	Id          int64   `json:"id"`
	OrderAmount float64 `json:"orderAmount"`
	OrderNumber string  `json:"orderNumber"`
	OrderTime   string  `json:"orderTime"`
}

type PaymentVO struct {
	EstimatedDeliveryTime string `json:"estimatedDeliveryTime"`
}

type UserPageVO struct {
	Total   int64         `json:"total"`
	Records []UserOrderVO `json:"records"`
}
type UserOrderVO struct {
	Id                    int64         `json:"id"`
	Number                string        `json:"number"`
	Status                int           `json:"status"`
	UserId                int64         `json:"userId"`
	AddressBookId         int64         `json:"addressBookId"`
	OrderTime             string        `json:"orderTime"`
	CheckoutTime          string        `json:"checkoutTime"`
	PayMethod             int           `json:"payMethod"`
	PayStatus             int           `json:"payStatus"`
	Amount                float64       `json:"amount"`
	Remark                string        `json:"remark"`
	Phone                 string        `json:"phone"`
	Address               string        `json:"address"`
	UserName              string        `json:"userName"`
	Consignee             string        `json:"consignee"`
	CancelReason          string        `json:"cancelReason"`
	RejectionReason       string        `json:"rejectionReason"`
	CancelTime            string        `json:"cancelTime"`
	EstimatedDeliveryTime string        `json:"estimatedDeliveryTime"`
	DeliveryStatus        int           `json:"deliveryStatus"`
	DeliveryTime          string        `json:"deliveryTime"`
	PackAmount            int           `json:"packAmount"`
	TableWareNumber       int           `json:"tablewareNumber"`
	TableWareStatus       int           `json:"tablewareStatus"`
	OrderDetails          []OrderDetail `json:"orderDetailList"`
}

type OrderVO struct {
	Address               string        `json:"address"`
	AddressBookId         int64         `json:"addressBookId"`
	Amount                float64       `json:"amount"`
	CancelReason          string        `json:"cancelReason"`
	CancelTime            string        `json:"cancelTime"`
	CheckoutTime          string        `json:"checkoutTime"`
	Consignee             string        `json:"consignee"`
	DeliveryStatus        int           `json:"deliveryStatus"`
	DeliveryTime          string        `json:"deliveryTime"`
	EstimatedDeliveryTime string        `json:"estimatedDeliveryTime"`
	Id                    int64         `json:"id"`
	Number                string        `json:"number"`
	OrderDetails          []OrderDetail `json:"orderDetailList"`
	OrderDishes           string        `json:"orderDishes"`
	OrderTime             string        `json:"orderTime"`
	PackAmount            int           `json:"packAmount"`
	PaymentMethod         int           `json:"paymentMethod"`
	PaymentStatus         int           `json:"paymentStatus"`
	Phone                 string        `json:"phone"`
	RejectionReason       string        `json:"rejectionReason"`
	Remark                string        `json:"remark"`
	Status                int           `json:"status"`
	TablewareNumber       int           `json:"tablewareNumber"`
	TablewareStatus       int           `json:"tablewareStatus"`
	UserId                int64         `json:"userId"`
	UserName              string        `json:"userName"`
}

type OrderDetail struct {
	Amount     float64 `json:"amount"`
	DishFlavor string  `json:"dishFlavor"`
	DishId     int64   `json:"dishId"`
	Id         int64   `json:"id"`
	Image      string  `json:"image"`
	Name       string  `json:"name"`
	Number     int     `json:"number"`
	OrderId    int64   `json:"orderId"`
	SetMealId  int64   `json:"setmealId"`
}

type AdminOrderVO struct {
	Id                    int64   `json:"id"`
	Number                string  `json:"number"`
	Status                int     `json:"status"`
	UserId                int64   `json:"userId"`
	AddressBookId         int64   `json:"addressBookId"`
	OrderTime             string  `json:"orderTime"`
	CheckoutTime          string  `json:"checkoutTime"`
	PayMethod             int     `json:"payMethod"`
	PayStatus             int     `json:"payStatus"`
	Amount                float64 `json:"amount"`
	Remark                string  `json:"remark"`
	Phone                 string  `json:"phone"`
	Address               string  `json:"address"`
	UserName              string  `json:"userName"`
	Consignee             string  `json:"consignee"`
	CancelReason          string  `json:"cancelReason"`
	RejectionReason       string  `json:"rejectionReason"`
	CancelTime            string  `json:"cancelTime"`
	EstimatedDeliveryTime string  `json:"estimatedDeliveryTime"`
	DeliveryStatus        int     `json:"deliveryStatus"`
	DeliveryTime          string  `json:"deliveryTime"`
	PackAmount            int     `json:"packAmount"`
	TableWareNumber       int     `json:"tablewareNumber"`
	TableWareStatus       int     `json:"tablewareStatus"`
	OrderDishes           string  `json:"orderDishes"`
}

type AdminPageVO struct {
	Total   int64          `json:"total"`
	Records []AdminOrderVO `json:"records"`
}

type StatisticsVO struct {
	Confirmed          int64 `json:"confirmed"`
	DeliveryInProgress int64 `json:"deliveryInProgress"`
	ToBeConfirmed      int64 `json:"toBeConfirmed"`
}
