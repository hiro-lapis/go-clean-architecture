package main

import (
	"clean-architecture/controller"
	"clean-architecture/db"
	"clean-architecture/repository"
	"clean-architecture/router"
	"clean-architecture/usecase"
	"clean-architecture/validator"
)

// entry point
func main() {
	// DB接続
	db := db.NewDB()
	// DB接続をリポジトリに渡す
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	// バリデーション構造体の初期化
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	// 各種InterfaceRepositoryを満たすリポジトリ実装をユースケースに渡す
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	// 各種IUsecaseを満たすユースケースをコントローラーに渡す
	userController := controller.NewUserConteroller(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	// 各種IControllerを満たすコントローラーをルーターに渡す
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
