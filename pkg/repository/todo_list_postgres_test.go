package repository

import (
	"errors"
	todo "github.com/LittleMikle/sber_it"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestTodoListPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal().Err(err).Msg("failed with sqlmock.Newx")
	}
	defer db.Close()

	r := NewTodoListPostgres(db)
	type args struct {
		input todo.TodoList
	}

	type mockBehavior func(args args, id int)
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				input: todo.TodoList{
					Title:       "Melushev",
					Description: "Mikhail",
					Date:        "2023.09.23",
					Status:      "undone",
				},
			},
			id: 0,
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO todo_lists").
					WithArgs(args.input.Title, args.input.Description, args.input.Date, args.input.Status).
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
		{
			name: "Empty Fields",
			args: args{
				input: todo.TodoList{
					Title: "",
				},
			},
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(1, errors.New("some error"))
				mock.ExpectQuery("INSERT INTO todo_lists").
					WithArgs(args.input.Title).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			got, err := r.Create(testCase.args.input)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
			}
		})
	}
}

func TestTodoListPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal().Err(err).Msg("failed with sqlmock.Newx")
	}
	defer db.Close()

	r := NewTodoListPostgres(db)

	type args struct {
		id int
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				id: 1,
			},
			mockBehavior: func(args args) {
				mock.ExpectExec("DELETE FROM todo_lists WHERE (.+)").
					WithArgs(args.id)
			},
			wantErr: true,
		},
		{
			name: "BAD QUERY",
			args: args{
				id: 1,
			},
			mockBehavior: func(args args) {

				mock.ExpectExec("INSERT INTO todo_lists").
					WithArgs(args.id)

			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			err := r.Delete(testCase.args.id)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
