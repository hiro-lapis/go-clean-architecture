package controller

import (
	"clean-architecture/model"
	"clean-architecture/usecase"
	"net/http"
	"os"
	"time"

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
	return &UserController{uu}
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

func (uc *UserController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// passing user variable, generate token
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// set token to current http cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	// JWTとは別にクッキー自体の有効期限は24hとする
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// ドメインのすべてのパスで有効なクッキーとする
	cookie.Path = "/"
	// ドメインを環境変数から取得し、セット
	cookie.Domain = os.Getenv("API_DOMAIN")
	// HTTPSのみで有効(テスト環境ではhttpで動作させるため、devの場合はfalse)
	if os.Getenv("APP_ENV") != "dev" {
		cookie.Secure = true
	}
	// クライアントサイドのスクリプトから読み取り不可
	cookie.HttpOnly = true
	// 異なるオリジンのリクエストにも送信
	cookie.SameSite = http.SameSiteNoneMode
	// cookieの内容をHTTPレスポンスヘッダにセットする
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *UserController) LogOut(e echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	// reset token value
	cookie.Value = ""
	// 初期化できればいいので、有効期限は現在日時
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	e.SetCookie(cookie)
	return e.NoContent(http.StatusOK)
}
