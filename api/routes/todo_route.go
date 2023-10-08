package routes

import (
	"cognixus/api/controllers"
	"cognixus/repository"
	todolist "cognixus/usecases/todo_list"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewTodoRoute(db *sqlx.DB, group *gin.RouterGroup) {
	repo := repository.New(db)
	uc := controllers.TodoListController{
		TodoListUsecase: todolist.New(repo),
	}
	group.GET("/todo", uc.GetTodoList)
	group.POST("/todo", uc.CreateTodo)
	group.PUT("/todo", uc.UpdateTodo)
	group.DELETE("/todo", uc.DeleteTodo)
}
