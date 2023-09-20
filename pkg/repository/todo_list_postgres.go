package repository

import (
	"fmt"
	todo "github.com/LittleMikle/sber_it"
	"github.com/jmoiron/sqlx"
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

func (r *TodoListPostgres) GetByStatus(params todo.TodoParams) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	getQuery := fmt.Sprintf("SELECT id, title, description, date, status FROM %s WHERE status=$1", todoListsTable)
	err := r.db.Select(&lists, getQuery, params.Status)
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

func (r *TodoListPostgres) Update(id int, updated todo.TodoList) error {
	return nil
}

func (r *TodoListPostgres) Delete(id int) error {
	return nil
}
