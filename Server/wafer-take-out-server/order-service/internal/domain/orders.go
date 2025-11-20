package domain

import "time"

type Order struct {
	Id                    int64      `gorm:"primaryKey;column:id;type:autoIncrement"`
	Number                string     `gorm:"column:number;Type:varchar(50)"`
	Status                int        `gorm:"column:status;Type:int"`
	UserId                int64      `gorm:"column:user_id;Type:bigint"`
	AddressBookId         int64      `gorm:"column:address_book_id;Type:bigint"`
	OrderTime             *time.Time `gorm:"column:order_time;Type:datetime"`
	CheckoutTime          *time.Time `gorm:"column:checkout_time;Type:datetime"`
	PayMethod             int        `gorm:"column:pay_method;Type:int,default:1"`
	PayStatus             int        `gorm:"column:pay_status;Type:tinyint"`
	Amount                float64    `gorm:"column:amount;Type:decimal(10,2)"`
	Remark                string     `gorm:"column:remark;Type:varchar(100)"`
	Phone                 string     `gorm:"column:phone;Type:varchar(11)"`
	Address               string     `gorm:"column:address;Type:varchar(255)"`
	UserName              string     `gorm:"column:user_name;Type:varchar(32)"`
	Consignee             string     `gorm:"column:consignee;Type:varchar(32)"`
	CancelReason          string     `gorm:"column:cancel_reason;Type:varchar(255)"`
	RejectionReason       string     `gorm:"column:rejection_reason;Type:varchar(255)"`
	CancelTime            *time.Time `gorm:"column:cancel_time;Type:datetime"`
	EstimatedDeliveryTime *time.Time `gorm:"column:estimated_delivery_time;Type:datetime"`
	DeliveryStatus        int        `gorm:"column:delivery_status;Type:tinyint(1);default:1"`
	DeliveryTime          *time.Time `gorm:"column:delivery_time;Type:datetime"`
	PackAmount            int        `gorm:"column:pack_amount;Type:int"`
	TableWareNumber       int        `gorm:"column:tableware_number;Type:int"`
	TableWareStatus       int        `gorm:"column:tableware_status;Type:tinyint(1);default:1"`
}

func (Order) TableName() string {
	return "orders"
}
