package validator

import (
	"clean-architecture/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

// TaskValidate is a function to validate task
// 引数は値渡し
func (uv *taskValidator) TaskValidate(task model.Task) error {
	// validation ライブラリを用いてrequest bindで作成したstructのバリデーション
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 strings"),
		))
}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}
