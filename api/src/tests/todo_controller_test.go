package tests

import (
	"bytes"
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"

	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/router"
)

type Todo entities.Todo
type TodoIndexRes entities.TodoIndexRes
type TodoShowRes entities.TodoShowRes
type TodoCreateRes entities.TodoCreateRes
type TodoUpdateRes entities.TodoUpdateRes

type TestTodoIndexRes struct {
	Todos []TodoIndexRes
}

type TestTodoShowRes struct {
	Todo TodoShowRes
}

type TestTodoCreateRes struct {
	Todo TodoCreateRes
}

type TestTodoUpdateRes struct {
	Todo TodoUpdateRes
}

func TestTodoIndex(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo, todo.ID)
	defer DeleteData(user, user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	b, _ := ioutil.ReadAll(w.Body)

	var res TestTodoIndexRes
	if err := json.Unmarshal(b, &res); err != nil {
		log.Println(err.Error())
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, res.Todos[0].ID, todo.ID)
	assert.Equal(t, res.Todos[0].Title, todo.Title)
	assert.Equal(t, res.Todos[0].Status, todo.Status)
	assert.Equal(t, res.Todos[0].UserID, todo.UserID)
	assert.Equal(t, res.Todos[0].User.ID, user.ID)
	assert.Equal(t, res.Todos[0].User.DisplayName, user.DisplayName)
}

func TestTodoShow(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo, todo.ID)
	defer DeleteData(user, user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	reqBody, _ := ioutil.ReadAll(w.Body)

	var res TestTodoShowRes
	if err := json.Unmarshal(reqBody, &res); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, res.Todo.ID, todo.ID)
	assert.Equal(t, res.Todo.UserID, user.ID)
	assert.Equal(t, res.Todo.User.ID, user.ID)
}

func TestTodoPost(t *testing.T) {
	user := CreateUser()
	defer DeleteData(user, user.ID)

	todo := new(Todo)
	todo.Title = "Test"
	todo_json, _ := json.Marshal(todo)
	body := bytes.NewBuffer(todo_json)

	r := router.Router()
	req, err := http.NewRequest("POST", "/todos", body)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	b, _ := ioutil.ReadAll(w.Body)

	var res TestTodoCreateRes
	if err := json.Unmarshal(b, &res); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, res.Todo.Title, "Test")
	assert.Equal(t, res.Todo.UserID, user.ID)
	assert.Equal(t, res.Todo.User.ID, user.ID)
}

func TestTodoUpdate(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo, todo.ID)
	defer DeleteData(user, user.ID)

	r := router.Router()
	req, err := http.NewRequest("PUT", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	b, _ := ioutil.ReadAll(w.Body)

	var res TestTodoUpdateRes
	if err := json.Unmarshal(b, &res); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, res.Todo.Status, true)
}

func TestTodoDelete(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo, todo.ID)
	defer DeleteData(user, user.ID)

	r := router.Router()
	req, err := http.NewRequest("DELETE", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
