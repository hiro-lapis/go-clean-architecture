package router

import (
	"clean-architecture/controller"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
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
	// CORS設定(origin,ヘッダ)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// アクセスを許可するドメイン
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		// 設定許可するHeader属性
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			//　ヘッダ経由でCSRFトークンを取得
			echo.HeaderXCSRFToken,
		},
		// 許可するHTTP method
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	// CSRF設定
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		// トークンはドメインのroot pathに設定
		CookiePath:   "/",
		CookieDomain: os.Getenv("API_DOMAIN"),
		// jsスクリプトからの読取拒否
		CookieHTTPOnly: true,
		// prod:SameSiteNoneMode(強制的にsecure=true) dev:SameSiteDefaultMode(postmanデバッグ)
		//CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge: 60,
	}))
	e.GET("/csrf", uc.CsrfToken)
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
