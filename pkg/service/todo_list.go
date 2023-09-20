package service

import (
	todo "github.com/LittleMikle/sber_it"
	"github.com/LittleMikle/sber_it/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(list todo.TodoList) (int, error) {
	return s.repo.Create(list)
}

func (s *TodoListService) GetLists(params todo.TodoParams) ([]todo.TodoList, error) {
	if params.Date != "" && params.Status == "" {
		return s.repo.GetByDate(params)
	}
	if params.Status != "" && params.Date == "" {
		return s.repo.GetByStatus(params)
	}
	if params.Date != "" && params.Status != "" {
		return s.repo.GetByParams(params)
	}
	return s.repo.GetLists()
}

func (s *TodoListService) Update(id int, updated todo.UpdateListInput) error {
	if err := updated.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, updated)
}

func (s *TodoListService) Delete(id int) error {
	return s.repo.Delete(id)
}
