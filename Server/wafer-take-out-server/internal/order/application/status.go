package application

import "time"

const (
	PENDING_PAYMENT      = 1
	TO_BE_CONFIRMED      = 2
	CONFIRMED            = 3
	DELIVERY_IN_PROGRESS = 4
	COMPLETED            = 5
	CANCELLED            = 6
)

const (
	UN_PAID = 0
	PAID    = 1
	REFUND  = 2
)

var MYSQL_MIN_TIME = time.Date(1000, 1, 1, 0, 0, 0, 0, time.Local)
