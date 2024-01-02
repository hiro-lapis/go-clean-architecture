package main

import (
	"clean-architecture/controller"
	"clean-architecture/db"
	"clean-architecture/repository"
	"clean-architecture/router"
	"clean-architecture/usecase"
)

// entry point
func main() {
	// DB接続
	db := db.NewDB()
	// DB接続をリポジトリに渡す
	userRepository := repository.NewUserRepository(db)
	// IUserRepositoryを満たすリポジトリをユースケースに渡す
	userUseCase := usecase.NewUserUseCase(userRepository)
	// IUserUsecaseを満たすユースケースをコントローラーに渡す
	userController := controller.NewUserConteroller(userUseCase)
	// IUserControllerを満たすコントローラーをルーターに渡す
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
