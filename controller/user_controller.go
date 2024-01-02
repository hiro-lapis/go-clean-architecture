package controller

import (
	"clean-architecture/usecase"

	"github.com/labstack/echo/v4"
)

// IUserController is an interface for user controller
// recieve request as echo.Context
type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

// UserController is a struct for user controller
// return is depend on usecase interface
type UserController struct {
	uu usecase.IUserUsecase
}

func NewUserConteroller(uu usecase.IUserUsecase) IUserController {
	return &UserController(uu)
}
