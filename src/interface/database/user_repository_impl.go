package database

import (
	"context"

	"github.com/takuma123-type/golang-study/src/domain/user"
	"golang.org/x/xerrors"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

var (
	u1, _ = user.GenForTest(user.NewUserID(), "user1", "lastuser1")
	u2, _ = user.GenForTest(user.NewUserID(), "user2", "lastuser2")
	u3, _ = user.GenForTest(user.NewUserID(), "user3", "lastuser3")
	users = []*user.User{
		u1,
		u2,
		u3,
	}
)

func (repo *UserRepositoryImpl) FindAll(ctx context.Context) ([]*user.User, error) {
	return users, nil
}

func (repo *UserRepositoryImpl) FindByID(ctx context.Context, id user.UserID) (*user.User, error) {
	for _, u := range users {
		if u.ID().Equal(id) {
			return u, nil
		}
	}
	return nil, xerrors.Errorf("userdm is not found: %s", id)
}
func (repo *UserRepositoryImpl) Store(ctx context.Context, user *user.User) error {
	// DBの保存書処理
	return nil
}
