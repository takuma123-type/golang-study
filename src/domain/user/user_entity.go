package user

import (
	"errors"
	"unicode/utf8"

	"github.com/takuma123-type/golang-study/src/domain/shared"
)

const (
	maxFirstNameLength = 30
	maxLastNameLength  = 30
)

type User struct {
	id        UserID
	firstName string
	lastName  string
	createdAt shared.CreatedAt
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) CreatedAt() shared.CreatedAt {
	return u.createdAt
}

func NewUser(id UserID, firstName string, lastName string) (*User, error) {
	if firstName == "" {
		return nil, errors.New("first name is required")
	}

	if lastName == "" {
		return nil, errors.New("last name is required")
	}

	if utf8.RuneCountInString(firstName) > maxFirstNameLength {
		return nil, errors.New("first name is too long (maximum length is 30)")
	}

	if utf8.RuneCountInString(lastName) > maxLastNameLength {
		return nil, errors.New("last name is too long (maximum length is 30)")
	}

	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		createdAt: shared.NewCreatedAt(),
	}, nil
}
