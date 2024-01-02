package repository

import (
	"clean-architecture/model"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(userId uint, taskId uint) error
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

type taskRepository struct {
	db *gorm.DB
}
