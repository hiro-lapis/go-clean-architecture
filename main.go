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
	taskRepository := repository.NewTaskRepository(db)
	// 各種InterfaceRepositoryを満たすリポジトリ実装をユースケースに渡す
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	// 各種IUsecaseを満たすユースケースをコントローラーに渡す
	userController := controller.NewUserConteroller(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	// 各種IControllerを満たすコントローラーをルーターに渡す
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
