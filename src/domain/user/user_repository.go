package user

import (
	"context"
)

type Repository interface {
	GetByID(ctx context.Context, id UserID) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, user *User) error
}
