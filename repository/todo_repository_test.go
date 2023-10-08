package repository_test

import (
	"cognixus/domain/entities"
	"cognixus/mocks"
	"cognixus/repository"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	tMock "github.com/stretchr/testify/mock"
)

const USER_ID = "2138930"

func TestCreate(t *testing.T) {
	var databaseHelper *mocks.TodoListRepository
	databaseHelper = &mocks.TodoListRepository{}

	mockTodo := entities.TodoList{
		ID:          uuid.New().String(),
		Name:        "Test",
		Description: "Test Description",
		Status:      false,
		UserID:      USER_ID,
	}

	t.Run("Get todo list from database with database error", func(t *testing.T) {
		db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		sqlxDB := sqlx.NewDb(db, "sqlmock")

		mockDB.ExpectQuery(`SELECT * FROM todo_items WHERE user_id=$1`).WithArgs(USER_ID).WillReturnError(errors.New("database error"))
		databaseHelper.On("Get", tMock.Anything).Return(nil, errors.New("database error")).Once()

		repo := repository.New(sqlxDB)
		items, err := repo.Get(context.Background(), USER_ID)

		assert.Equal(t, errors.New("database error"), err)
		assert.Nil(t, items)
	})

	t.Run("Create todo item into database with database error", func(t *testing.T) {
		db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		sqlxDB := sqlx.NewDb(db, "sqlmock")

		mockDB.ExpectQuery(`INSERT INTO todo_items (id, user_id, name, description, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`).WithArgs(mockTodo.ID, mockTodo.UserID, mockTodo.Name, mockTodo.Description, mockTodo.Status).WillReturnError(errors.New("database error"))
		databaseHelper.On("Create", tMock.Anything).Return(nil, errors.New("database error")).Once()

		repo := repository.New(sqlxDB)
		err = repo.Create(context.Background(), &mockTodo)

		assert.Equal(t, errors.New("database error"), err)
	})
}
