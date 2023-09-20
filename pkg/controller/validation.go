package controller

import (
	"errors"
	todo "github.com/LittleMikle/sber_it"
)

func validateInput(input todo.TodoList) error {
	if input.Title == "" {
		return errors.New("title can't be empty")
	}
	if input.Date == "" {
		return errors.New("date can't be empty")
	}
	return nil
}
