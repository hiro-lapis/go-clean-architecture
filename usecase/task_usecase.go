package usecase

import (
	"clean-architecture/model"
	"clean-architecture/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId, taskId uint) error
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

type taskUsecase struct {
	tr repository.ITaskRepository
}
