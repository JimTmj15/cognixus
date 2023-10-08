package entities

import (
	"context"
)

type TodoList struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Status      bool   `db:"status"`
	UserID      string `db:"user_id"`
}

type TodoListRepository interface {
	Create(c context.Context, todoList *TodoList) error
	Update(c context.Context, id string, userId string) error
	Delete(c context.Context, id string, userId string) error
	Get(c context.Context, userId string) (*[]TodoList, error)
}

type TodoListUsecase interface {
	CreateTodo(c context.Context, todoList *TodoList) error
	UpdateTodo(c context.Context, id string, userId string) error
	DeleteTodo(c context.Context, id string, userId string) error
	GetTodoList(c context.Context, userId string) (*[]TodoList, error)
}
