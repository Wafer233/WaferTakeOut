package domain

import "time"

type OrderDomainService struct {
}

func NewOrderDomainService() *OrderDomainService {
	return &OrderDomainService{}
}

func (svc *OrderDomainService) ParseTime(tm *time.Time) string {

	layout := "2006-01-02 15:04:05"

	if tm == nil {
		return ""
	}

	return tm.Format(layout)
}
