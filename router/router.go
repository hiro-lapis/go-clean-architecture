package router

import (
	"clean-architecture/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// request handler
// リクエストを受け取り、コントローラーに処理を委譲する
func NewRouter(
	uc controller.IUserController,
	tc controller.ITaskController,
) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)

	// route group作成
	t := e.Group("/tasks")
	// task関連の処理にmiddlewareを設定
	t.Use(echojwt.WithConfig(echojwt.Config{
		// JWT生成時と同じ鍵
		SigningKey: []byte(os.Getenv("SECRET")),
		// token参照先
		TokenLookup: "cookie:token",
	}))
	// group path以降のURLでrouting設定
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	// ハンドラ設定をしたechoインスタンスを返す
	return e
}
