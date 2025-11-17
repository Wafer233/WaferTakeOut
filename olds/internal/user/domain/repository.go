package domain

import "context"

type UserRepository interface {
	Upsert(context.Context, *User) error
}
