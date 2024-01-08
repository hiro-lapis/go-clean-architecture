package repository

import (
	"clean-architecture/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

type taskRepository struct {
	db *gorm.DB
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	// 書き方のバリエーション
	// pattern1 Firstの引数条件を渡す
	if err := tr.db.Where("user_id = ?)", userId).First(task, taskId).Error; err != nil {
		return err
	}
	// pattern2 条件の数だけWhereを書く
	// if err := tr.db.Where("user_id = ?)", userId).Where("id = ?", taskId).First(task).Error; err != nil {
	// 	return err
	// }
	// pattern3 Whereは1つで条件を複数渡す
	// if err := tr.db.Where("user_id = ? AND id = ?", userId, taskId).First(task).Error; err != nil {
	// 	return err
	// }
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("data doesn't exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id = ? AND user_id = ?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("data doesn't exist")
	}
	return nil
}
