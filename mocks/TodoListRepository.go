// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	entities "cognixus/domain/entities"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// TodoListRepository is an autogenerated mock type for the TodoListRepository type
type TodoListRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, todoList
func (_m *TodoListRepository) Create(c context.Context, todoList *entities.TodoList) error {
	ret := _m.Called(c, todoList)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.TodoList) error); ok {
		r0 = rf(c, todoList)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: c, id, userId
func (_m *TodoListRepository) Delete(c context.Context, id string, userId string) error {
	ret := _m.Called(c, id, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(c, id, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: c, userId
func (_m *TodoListRepository) Get(c context.Context, userId string) (*[]entities.TodoList, error) {
	ret := _m.Called(c, userId)

	var r0 *[]entities.TodoList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*[]entities.TodoList, error)); ok {
		return rf(c, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *[]entities.TodoList); ok {
		r0 = rf(c, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.TodoList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, id, userId
func (_m *TodoListRepository) Update(c context.Context, id string, userId string) error {
	ret := _m.Called(c, id, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(c, id, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTodoListRepository creates a new instance of TodoListRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTodoListRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TodoListRepository {
	mock := &TodoListRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
