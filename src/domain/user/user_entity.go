package user

import (
	"errors"

	"github.com/hrs-o/docker-go/domain/shared"
)

const (
	maxFirstNameLength = 30
	maxLastNameLength  = 30
)

type User struct {
	ID        UserID
	FirstName string
	LastName  string
	CreatedAt shared.CreatedAt
}

func (u *User) GetCreatedAt() shared.CreatedAt {
	return u.CreatedAt
}

func NewUser(id UserID, firstName string, lastName string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("first name is required")
	}

	if lastName == "" {
		return nil, errors.New("last name is required")
	}

	if len(firstName) > maxFirstNameLength {
		return nil, errors.New("first name is too long (maximum length is 30)")
	}

	if len(lastName) > maxLastNameLength {
		return nil, errors.New("last name is too long (maximum length is 30)")
	}

	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: shared.NewCreatedAt(),
	}, nil
}
