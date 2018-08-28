package controller

import (
	"commenting/interface/presenter"
	"commenting/interface/repository"
	"commenting/usecase"
	"github.com/labstack/echo"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (ctl *UserController) CurrentUser(c echo.Context) error {
	user, result := usecase.NewUserUsecase(
		repository.NewUserRepository(c.StdContext()),
	).GetCurrentUser()
	return presenter.RenderJSON(c, user, result)
}
