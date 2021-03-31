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

type TodoPost struct {
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
	r := router.Router()
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		log.Println(err.Error())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestTodoShow(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)

	r := router.Router()

	req, err := http.NewRequest("GET", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	DeleteData(user.ID)
}

func TestTodoPost(t *testing.T) {
	user := CreateUser()

	r := router.Router()

	title := "Test"
	expectedTodo := entities.Todo{
		Title: title,
	}

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

	var resData TodoPost
	if err := json.Unmarshal(b, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, resData.Todo.Title, expectedTodo.Title)

	DeleteData(user.ID)
}

func TestTodoUpdate(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)

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

	var resData TodoPost
	if err := json.Unmarshal(b, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.Todo.Status, true)

	DeleteData(user.ID)
}

func TestTodoDelete(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)

	r := router.Router()

	req, err := http.NewRequest("DELETE", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)

	DeleteData(user.ID)
}
