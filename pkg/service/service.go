package service

import (
	todo "github.com/LittleMikle/sber_it"
	"github.com/LittleMikle/sber_it/pkg/repository"
)

type TodoList interface {
	Create(list todo.TodoList) (int, error)
	GetLists(params todo.TodoParams) ([]todo.TodoList, error)
	Update(id int, updated todo.TodoList) error
	Delete(id int) error
}

type Service struct {
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(repos.TodoList),
	}
}
