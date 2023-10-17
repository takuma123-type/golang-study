package userusecase

import (
	"context"

	"github.com/takuma123-type/golang-study/src/domain/user"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase/userinput"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase/useroutput"
)

type GetUserByIDUsecase struct {
	userRepository user.UserRepository
}

func NewGetUserByID(userRepo user.UserRepository) *GetUserByIDUsecase {
	return &GetUserByIDUsecase{
		userRepository: userRepo,
	}
}

func (use *GetUserByIDUsecase) Exec(ctx context.Context, in *userinput.GetUserByIDInput) (*useroutput.UserByID, error) {
	userID, err := user.NewUserIDByVal(in.ID)
	if err != nil {
		return nil, err
	}
	user, err := use.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &useroutput.UserByID{
		ID:        user.ID().String(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
	}, nil
}
