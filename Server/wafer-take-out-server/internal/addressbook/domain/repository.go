package domain

import "context"

type AddressRepository interface {
	Create(context.Context, *AddressBook) error
	FindByUserId(context.Context, int64) ([]*AddressBook, error)
}
