package controllers

import (
	"cognixus/domain/entities"
	"cognixus/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoListController struct {
	TodoListUsecase entities.TodoListUsecase
}

func (t *TodoListController) CreateTodo(c *gin.Context) {
	var req models.CreateTodoRequest
	uid := c.Request.Header.Get("X-USER-ID")

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	todo := entities.TodoList{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Status:      false,
		UserID:      uid,
	}

	todoErr := t.TodoListUsecase.CreateTodo(c, &todo)
	if todoErr != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: todoErr.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (t *TodoListController) GetTodoList(c *gin.Context) {
	var request models.GetTodoRequest
	uid := c.Request.Header.Get("X-USER-ID")

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	todoList, err := t.TodoListUsecase.GetTodoList(c, uid)

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Get todo list successfully",
		Data:    todoList,
	})
}

func (t *TodoListController) UpdateTodo(c *gin.Context) {
	var req models.TodoIdRequest
	uid := c.Request.Header.Get("X-USER-ID")

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	dbErr := t.TodoListUsecase.UpdateTodo(c, req.ID, uid)
	if dbErr != nil {
		c.JSON(http.StatusOK, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Updated todo successfully",
		Data:    nil,
	})
}

func (t *TodoListController) DeleteTodo(c *gin.Context) {
	var req models.TodoIdRequest
	uid := c.Request.Header.Get("X-USER-ID")

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	dbErr := t.TodoListUsecase.DeleteTodo(c, req.ID, uid)
	if dbErr != nil {
		c.JSON(http.StatusOK, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Delete todo successfully",
		Data:    nil,
	})
}
