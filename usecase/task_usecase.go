package usecase

import (
	"clean-architecture/model"
	"clean-architecture/repository"
	"clean-architecture/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId, taskId uint) error
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	// 1.prepare slice of task model
	tasks := []model.Task{}
	tastRes := []model.TaskResponse{}
	// 2.get all tasks by user id
	// tasksのポインタを渡し、取得した情報で書き換えてもらう
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return tastRes, err
	}
	// 3.convert task model to task response model
	for _, task := range tasks {
		tastRes = append(tastRes, model.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}
	return tastRes, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	response := model.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	return response, nil
}

func (tu *taskUsecase) GetTaskById(userId, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	response := model.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	return response, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userId, taskId uint) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	response := model.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	return response, nil
}

func (tu *taskUsecase) DeleteTask(userId, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	// 正常削除時はnilを返す
	return nil
}
