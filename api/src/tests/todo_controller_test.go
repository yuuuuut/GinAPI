package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"

	"strings"
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/router"
	"github.com/yuuuuut/gin-api/src/util"
)

type Todo entities.Todo

type IndexTodo struct {
	Todos []Todo
}

type OneTodo struct {
	Todo struct {
		ID     int
		Title  string
		Status bool
		UserID string
		User   entities.User
	}
}

func TestMain(m *testing.M) {
	util.InitTestENV()
	util.InitTestDB()
	util.InitTestFirebase()

	code := m.Run()

	os.Exit(code)
}

func TestTodoIndex(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo.ID)
	defer DeleteData(user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	reqBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData IndexTodo
	if err := json.Unmarshal(reqBody, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.Todos[0].ID, todo.ID)
	assert.Equal(t, resData.Todos[0].UserID, todo.UserID)
	assert.Equal(t, resData.Todos[0].User, user)
}

func TestTodoShow(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo.ID)
	defer DeleteData(user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	reqBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData OneTodo
	if err := json.Unmarshal(reqBody, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.Todo.ID, todo.ID)
	assert.Equal(t, resData.Todo.UserID, user.ID)
	assert.Equal(t, resData.Todo.User, user)
}

func TestTodoPost(t *testing.T) {
	user := CreateUser()
	defer DeleteData(user.ID)

	title := "Test"

	r := router.Router()
	body := strings.NewReader(fmt.Sprintf(`{"title": "%s"}`, title))
	req, err := http.NewRequest("POST", "/todos", body)
	if err != nil {
		log.Println(err.Error())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData OneTodo
	if err := json.Unmarshal(b, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, resData.Todo.Title, title)
	assert.Equal(t, resData.Todo.User, user)
}

func TestTodoUpdate(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo.ID)
	defer DeleteData(user.ID)

	r := router.Router()
	req, err := http.NewRequest("PUT", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData OneTodo
	if err := json.Unmarshal(b, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.Todo.Status, true)
}

func TestTodoDelete(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo.ID)
	defer DeleteData(user.ID)

	r := router.Router()
	req, err := http.NewRequest("DELETE", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}
