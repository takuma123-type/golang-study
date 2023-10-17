package presenter

import (
	"context"
	"net/http"

	"github.com/takuma123-type/golang-study/src/domain/user"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase/userinput"
	"github.com/takuma123-type/golang-study/src/usecase/userusecase/useroutput"
)

type userPresent struct {
	delivery Presenter
}

func NewUserPresenter(p Presenter) UserPresenter {
	return &userPresent{
		delivery: p,
	}
}

type UserPresenter interface {
	UserList(users []*user.User)
	User(user *user.User)
	Create(out *useroutput.CreateUserOutput)
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

func (p *userPresent) UserList(out []*user.User) {
	p.delivery.JSON(http.StatusOK, out)
}

func (p *userPresent) User(out *user.User) {
	p.delivery.JSON(http.StatusOK, out)
}

func (p *userPresent) Create(out *useroutput.CreateUserOutput) {
	p.delivery.JSON(http.StatusCreated, out)
}
