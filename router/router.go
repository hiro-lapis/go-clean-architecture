package router

import (
	"clean-architecture/controller"

	"github.com/labstack/echo/v4"
)

// request handler
// リクエストを受け取り、コントローラーに処理を委譲する
func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)
	// ハンドラ設定をしたechoインスタンスを返す
	return e
}
