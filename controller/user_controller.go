package controller

import (
	"clean-architecture/model"
	"clean-architecture/usecase"
	"net/http"

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

func (uc *UserController) SignUp(c echo.Context) error {
	// 1.prepare variable for bind request context
	user := model.User{}
	// 2. bind variable with request context
	// リクエストヘッダのcontent typeに応じて入力値を参照し、変数に値を注入する
	if err := c.Bind(&user); err != nil {
		// HTTP Status 400 Bad Request, error message
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// 3. call usecase, create user
	// by passing binded variable, return response created by user usecase
	// ユースケースはコントローラのプロパティとなっていて、それを呼び出して処理を行う
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		// HTTP Status 500 Internal Server Error, error message
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 4. set response info to context
	return c.JSON(http.StatusCreated, userRes)
}

func NewUserConteroller(uu usecase.IUserUsecase) IUserController {
	return &UserController{uu}
}
