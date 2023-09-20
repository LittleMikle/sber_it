package todo

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Date        string `json:"date" db:"date"`
	Status      string `json:"status" db:"status"`
}

type TodoParams struct {
	Date   string
	Status string
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Date        *string `json:"date"`
	Status      *string `json:"status"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Date == nil && i.Status == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
