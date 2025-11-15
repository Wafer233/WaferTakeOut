package domain

import "context"

type AddressRepository interface {
	Create(context.Context, *AddressBook) error
	FindByUserId(context.Context, int64) ([]*AddressBook, error)
	FindByUserIdDefault(context.Context, int64) (*AddressBook, error)
	UpdateDefault(context.Context, int64, int64, int) error
	FindById(context.Context, int64) (*AddressBook, error)
	DeleteById(context.Context, int64) error
	Update(context.Context, *AddressBook) error
}
