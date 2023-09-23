package controller

import (
	"bytes"
	todo "github.com/LittleMikle/sber_it"
	"github.com/LittleMikle/sber_it/pkg/service"
	mock_service "github.com/LittleMikle/sber_it/pkg/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandler_createList(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTodoList, todoList todo.TodoList)

	testTable := []struct {
		name                string
		inputBody           string
		inputToDo           todo.TodoList
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"title":"Melushev", "description":"Mikhail", "date":"2023.09.23"}`,
			inputToDo: todo.TodoList{
				Title:       "Melushev",
				Description: "Mikhail",
				Date:        "2023.09.23",
			},
			mockBehavior: func(s *mock_service.MockTodoList, todoList todo.TodoList) {
				s.EXPECT().Create(todoList).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "No title",
			inputBody: `{"title":"", "description":"Mikhail", "date":"2023.09.23"}`,
			inputToDo: todo.TodoList{
				Title:       "",
				Description: "Mikhail",
				Date:        "2023.09.23",
			},
			mockBehavior:        func(s *mock_service.MockTodoList, todoList todo.TodoList) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"title can't be empty"}`,
		},
		{
			name:      "No date",
			inputBody: `{"title":"Melushev", "description":"Mikhail", "date":""}`,
			inputToDo: todo.TodoList{
				Title:       "Melushev",
				Description: "Mikhail",
				Date:        "",
			},
			mockBehavior:        func(s *mock_service.MockTodoList, todoList todo.TodoList) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"date can't be empty"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			create := mock_service.NewMockTodoList(c)
			testCase.mockBehavior(create, testCase.inputToDo)

			services := &service.Service{
				TodoList: create,
			}
			handler := NewController(services)

			r := gin.New()
			r.POST("/create", handler.createList)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_deleteLists(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTodoList, id int)

	testTable := []struct {
		name                 string
		id                   int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			id:   1,
			mockBehavior: func(s *mock_service.MockTodoList, id int) {
				s.EXPECT().Delete(id).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":"ok"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			del := mock_service.NewMockTodoList(c)
			testCase.mockBehavior(del, testCase.id)

			services := &service.Service{
				TodoList: del,
			}
			handler := NewController(services)
			r := gin.New()

			r.DELETE("/:id", handler.deleteList)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/"+strconv.Itoa(testCase.id), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}

func TestHandler_getLists(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTodoList, params todo.TodoParams)

	testTable := []struct {
		name                 string
		params               string
		inputParams          todo.TodoParams
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "OK date",
			params: "2023.09.23",
			inputParams: todo.TodoParams{
				Date: "2023.09.23",
			},
			mockBehavior: func(s *mock_service.MockTodoList, params todo.TodoParams) {
				s.EXPECT().GetLists(0, params).Return([]todo.TodoList{}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":[]}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			get := mock_service.NewMockTodoList(c)
			testCase.mockBehavior(get, testCase.inputParams)

			services := &service.Service{
				TodoList: get,
			}
			handler := NewController(services)

			r := gin.New()
			r.GET("/lists", handler.getLists)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/lists?date="+testCase.params, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
