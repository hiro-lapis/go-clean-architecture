package repository

import "clean-architecture/model"

// DBとのやりとりを介したCRUD処理のinterface定義をしているパッケージ
// usecaseから呼び出され、実装している実structとの仲介を行う
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
