package service

import (
	todo "github.com/LittleMikle/sber_it"
	"github.com/LittleMikle/sber_it/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go

type TodoList interface {
	Create(list todo.TodoList) (int, error)
	GetLists(page int, params todo.TodoParams) ([]todo.TodoList, error)
	Update(id int, updated todo.UpdateListInput) error
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
