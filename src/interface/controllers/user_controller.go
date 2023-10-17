package controllers

import (
	"context"

	userinput "github.com/abyssparanoia/catharsis/user/interface/input/user"
	presenter "github.com/abyssparanoia/catharsis/user/interface/presenter/user"
	userusecase "github.com/abyssparanoia/catharsis/user/usecase/user"
	userdm "github.com/revenue-hack/cleanarchitecture-sample/src/domain/userdm"
)

type userController struct {
	delivery presenter.UserPresenter
	userRepo userdm.UserRepository
}

func NewUserController(p presenter.UserPresenter, userRepo userdm.UserRepository) *userController {
	return &userController{
		delivery: p,
		userRepo: userRepo,
	}
}

func (c *userController) GetUserList(ctx context.Context) error {
	usecase := userusecase.NewGetUserList(c.userRepo)
	out, err := usecase.Exec(ctx)
	if err != nil {
		return err
	}
	c.delivery.UserList(out)
	return nil

}

func (c *userController) GetUserByID(ctx context.Context, in *userinput.GetUserByIDInput) error {
	usecase := userusecase.NewGetUserByID(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.User(out)
	return nil
}
