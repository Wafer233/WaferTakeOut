package random

import (
	"math/rand"
	"time"
)

func GenerateOrderID() string {
	rand.Seed(time.Now().UnixNano())

	// 你可以自由扩展这个字符池
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*"

	length := 10
	orderID := make([]byte, length)

	for i := 0; i < length; i++ {
		orderID[i] = charset[rand.Intn(len(charset))]
	}

	return string(orderID)
}
