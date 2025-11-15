package domain

import "context"

type AddressRepository interface {
	Create(context.Context, *AddressBook) error
}
