package usecase

import (
	"clean-architecture/model"
	"clean-architecture/repository"
	"clean-architecture/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const EXP_HOUR = 12

// IUserUsecase is an interface that defines the user related logic and realizes the actual processing through the domain via the repository and other structs.
type IUserUsecase interface {
	// 入力(email, password)を元にユーザ情報を作成、ログイン処理を行う
	SignUp(user model.User) (model.UserResponse, error)
	// 入力(email, password)を元にユーザ情報を照合,ログイン処理を行い,JWTトークンを返す
	Login(user model.User) (string, error)
}

// userUsecase is a struct that implements the IUserUsecase interface.
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// 1. convert plain password to hash
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	// 2. create new user by using model, and pass it to repository
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// 3. create response
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	// 1. for assign check target user, prepare user model
	storedUser := model.User{}
	// 2. get user by email
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// 3. compare password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	// 4. create token value  with userID, exp
	// claims:トークン内の情報種別
	// Registered Claims(予約済): exp, iat, nbf, jtiなど
	// Public Claims(公開): iss, aud, subなど
	// Private Claims(非公開): その他
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * EXP_HOUR).Unix(),
	})
	// 5. login signing token
	tokenStrign, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenStrign, nil
}

// NewUserUsecase is a function that returns a userUsecase struct.
// interfaceを実装したリポジトリを引数にとり,interfaceを実装したusercaseの実クラスを返す関数
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}
