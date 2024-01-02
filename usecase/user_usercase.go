package usecase

import (
	"clean-architecture/model"
	"clean-architecture/repository"
)

// IUserUsecase is an interface that defines the user related logic and realizes the actual processing through the domain via the repository and other structs.
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (model.UserResponse, error)
}

// userUsecase is a struct that implements the IUserUsecase interface.
type userUsecase struct {
	ur repository.IUserRepository
}

// NewUserUseCase is a function that returns a userUsecase struct.
// interfaceを実装したリポジトリを引数にとり,interfaceを実装したusercaseの実クラスを返す関数
func NewUserUseCase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}
