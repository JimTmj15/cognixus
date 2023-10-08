package todolist

import (
	"cognixus/domain/entities"
	"context"
	"time"
)

type todoListUsecase struct {
	repo entities.TodoListRepository
}

func New(repo entities.TodoListRepository) entities.TodoListUsecase {
	return &todoListUsecase{
		repo: repo,
	}
}

func (u *todoListUsecase) CreateTodo(c context.Context, todoList *entities.TodoList) error {
	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()
	return u.repo.Create(ctx, todoList)
}

func (u *todoListUsecase) UpdateTodo(c context.Context, id string, userId string) error {
	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()
	return u.repo.Update(ctx, id, userId)
}

func (u *todoListUsecase) DeleteTodo(c context.Context, id string, userId string) error {
	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()
	return u.repo.Delete(ctx, id, userId)
}

func (u *todoListUsecase) GetTodoList(c context.Context, userId string) (*[]entities.TodoList, error) {
	ctx, cancel := context.WithTimeout(c, 2*time.Second)
	defer cancel()
	return u.repo.Get(ctx, userId)
}
