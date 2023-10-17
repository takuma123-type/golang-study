package controllers

import (
	"context"

	"github.com/takuma123-type/golang-study/src/domain/user"
	"github.com/takuma123-type/golang-study/src/interface/presenter"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase/userinput"
)

type userController struct {
	userRepo user.Repository
	delivery presenter.UserPresenter
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

func NewUserController(p presenter.UserPresenter, userRepo user.UserRepository) *userController {
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

func (c *userController) CreateUser(ctx context.Context, in *userinput.CreateUserInput) error {
	usecase := userusecase.NewCreateUser(c.userRepo)
	out, err := usecase.Exec(ctx, in)
	if err != nil {
		return err
	}
	c.delivery.Create(out)
	return nil
}
