package repository

import (
	"fmt"
	todo "github.com/LittleMikle/sber_it"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (title, description, date, status) VALUES ($1, $2, $3, $4) RETURNING id", todoListsTable)
	list.Status = "undone"
	row := tx.QueryRow(createQuery, list.Title, list.Description, list.Date, list.Status)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetLists() ([]todo.TodoList, error) {
	var lists []todo.TodoList
	getQuery := fmt.Sprintf("SELECT id, title, description, date, status FROM %s", todoListsTable)
	err := r.db.Select(&lists, getQuery)
	if err != nil {
		return lists, fmt.Errorf("failed with get lists")
	}
	return lists, nil
}

func (r *TodoListPostgres) GetByDate(params todo.TodoParams) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	getQuery := fmt.Sprintf("SELECT id, title, description, date, status FROM %s WHERE date=$1", todoListsTable)
	err := r.db.Select(&lists, getQuery, params.Date)
	if err != nil {
		return lists, fmt.Errorf("failed with get lists by date")
	}
	return lists, nil
}

func (r *TodoListPostgres) GetByStatus(page int, params todo.TodoParams) ([]todo.TodoList, error) {
	var offset = (page - 1) * 3
	var lists []todo.TodoList
	getQuery := fmt.Sprintf("SELECT id, title, description, date, status FROM %s WHERE status =$1 ORDER BY id OFFSET $2 LIMIT 3", todoListsTable)
	err := r.db.Select(&lists, getQuery, params.Status, offset)
	if err != nil {
		return lists, fmt.Errorf("failed with get lists by status")
	}
	return lists, nil
}

func (r *TodoListPostgres) GetByParams(params todo.TodoParams) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	getQuery := fmt.Sprintf("SELECT id, title, description, date, status FROM %s WHERE date=$1 AND status=$2", todoListsTable)
	err := r.db.Select(&lists, getQuery, params.Date, params.Status)
	if err != nil {
		return lists, fmt.Errorf("failed with get lists by status and date")
	}
	return lists, nil
}

func (r *TodoListPostgres) Update(id int, updated todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updated.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *updated.Title)
		argId++
	}

	if updated.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *updated.Description)
		argId++
	}

	if updated.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, *updated.Date)
		argId++
	}

	if updated.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *updated.Status)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	updateQuery := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d",
		todoListsTable, setQuery, id)
	_, err := r.db.Exec(updateQuery, args...)
	return err
}

func (r *TodoListPostgres) Delete(id int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s  WHERE id=$1",
		todoListsTable)
	_, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}
