package userusecase

import (
	"context"

	"github.com/takuma123-type/golang-study/src/domain/user"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase/useroutput"
)

type GetUserListUsecase struct {
	userRepository user.UserRepository
}

func NewGetUserList(userRepo user.UserRepository) *GetUserListUsecase {
	return &GetUserListUsecase{
		userRepository: userRepo,
	}
}

func (use *GetUserListUsecase) Exec(ctx context.Context) (*useroutput.UserList, error) {
	users, err := use.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	userItems := make([]*useroutput.UserItem, len(users))
	for i, user := range users {
		userItems[i] = &useroutput.UserItem{
			ID:        user.ID().String(),
			FirstName: user.FirstName(),
			LastName:  user.LastName(),
		}
	}
	return &useroutput.UserList{Users: userItems}, nil
}
