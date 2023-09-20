package repository

import (
	todo "github.com/LittleMikle/sber_it"
	"github.com/jmoiron/sqlx"
)

type TodoList interface {
	Create(list todo.TodoList) (int, error)
	GetLists() ([]todo.TodoList, error)
	GetByDate(params todo.TodoParams) ([]todo.TodoList, error)
	GetByStatus(params todo.TodoParams) ([]todo.TodoList, error)
	GetByParams(params todo.TodoParams) ([]todo.TodoList, error)
	Update(id int, updated todo.UpdateListInput) error
	Delete(id int) error
}

type Repository struct {
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoList: NewTodoListPostgres(db),
	}
}
