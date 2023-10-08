package models

type GetTodoRequest struct {
	ID string
}

type CreateTodoRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TodoIdRequest struct {
	ID string `json:"id" binding:"required"`
}
