package controller

import (
	"clean-architecture/model"
	"clean-architecture/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	// JWT tokenを取得
	// ref: https://pkg.go.dev/github.com/golang-jwt/jwt/v4@v4.5.0#Token
	user := c.Get("user").(*jwt.Token)
	// token内のclaimsをmaps(連装配列)で取得
	claims := user.Claims.(jwt.MapClaims)
	// token内のuser_idを取得
	userId := claims["user_id"]

	// any型->float64 型assertion->uint型にcast
	taskRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(echo.ErrHTTPVersionNotSupported.Code, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	// URLパスパラメータtask_idの値を取得
	id := c.Param("task_id")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	// 入力値からtaskモデル作成
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザ情報を取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	if userId == nil {
		return c.JSON(http.StatusBadRequest, "user_id is empty")
	}
	// taskに作成者情報をセット
	task.UserID = uint(userId.(float64))
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	// 入力値からtaskモデル作成
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザ情報を取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	if userId == nil {
		return c.JSON(http.StatusBadRequest, "user_id is empty")
	}
	id := c.Param("task_id")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	if userId == nil {
		return c.JSON(http.StatusBadRequest, "user_id is empty")
	}

	id := c.Param("task_id")
	taskId, _ := strconv.Atoi(id)
	if err := tc.tu.DeleteTask(uint(taskId), uint(userId.(float64))); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
