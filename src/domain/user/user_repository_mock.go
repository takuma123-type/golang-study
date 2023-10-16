package user

import "context"

type MockUserRepository struct {
	StoreErr     error
	FindAllUsers []*User
	FindAllErr   error
	FindByIDUser *User
	FindByIDErr  error
}

func (m *MockUserRepository) Store(ctx context.Context, user *User) error {
	return m.StoreErr
}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]*User, error) {
	return m.FindAllUsers, m.FindAllErr
}

func (m *MockUserRepository) FindByID(ctx context.Context, userID UserID) (*User, error) {
	return m.FindByIDUser, m.FindByIDErr
}
