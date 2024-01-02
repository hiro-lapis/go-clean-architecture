package repository

import (
	"clean-architecture/model"

	"gorm.io/gorm"
)

// DBとのやりとりを介したCRUD処理のinterface定義をしているパッケージ
// usecaseから呼び出され、実装している実structとの仲介を行う
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// userRepository is a struct that implements the IUserRepository interface.
type userRepository struct {
	db *gorm.DB
}

// constructor
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository(db)
}

// GetByEmail is a method that gets a user by email.
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}
