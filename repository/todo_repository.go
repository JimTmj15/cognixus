package repository

import (
	"cognixus/domain/entities"
	"context"

	"github.com/jmoiron/sqlx"
)

type todoRespository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) entities.TodoListRepository {
	return &todoRespository{
		db: db,
	}
}

func (r *todoRespository) Create(c context.Context, todo *entities.TodoList) error {
	sql := `INSERT INTO todo_items (id, user_id, name, description, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	row, err := r.db.QueryContext(c, sql, todo.ID, todo.UserID, todo.Name, todo.Description, todo.Status)
	if err != nil {
		return err
	}

	defer row.Close()

	var id string
	row.Next()

	err = row.Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *todoRespository) Update(c context.Context, id, userId string) error {
	tx := r.db.MustBegin()
	tx.MustExec("UPDATE todo_items SET status=$1 WHERE id=$2 AND user_id=$3", true, id, userId)
	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (r *todoRespository) Delete(c context.Context, id, userId string) error {
	tx := r.db.MustBegin()
	tx.MustExec("DELETE FROM todo_items WHERE id=$1 AND user_id=$2", id, userId)
	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (r *todoRespository) Get(c context.Context, userId string) (*[]entities.TodoList, error) {
	cons := new([]entities.TodoList)

	err := r.db.SelectContext(c, cons, "SELECT * FROM todo_items WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}

	return cons, nil
}
