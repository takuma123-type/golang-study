package controller

import (
	"context"

	"github.com/takuma123-type/golang-study/src/domain/userdm"
	"github.com/takuma123-type/golang-study/src/interface/presenter"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase/userinput"
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
	c.delivery.UserByID(out)
	return nil
}

func (c *userController) CreateUser(ctx context.Context, in *userinput.CreateUserInput) error {
	usecase := userusecase.NewCreateUser(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.Create(out)
	return nil
}
