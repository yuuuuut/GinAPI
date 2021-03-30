package tests

import (
	"bytes"
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

func UnmarshalData(body *bytes.Buffer) TodoPost {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	var todo TodoPost
	if err := json.Unmarshal(b, &todo); err != nil {
		log.Fatal(err)
	}

	return todo
}

func TestTodoIndex(t *testing.T) {
	var db = util.GetDB()
	todo := CreateTodo()

	r := router.Router()

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		log.Println(err.Error())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var todos []Todo
	db.Find(&todos)

	assert.Equal(t, len(todos), 1)

	DeleteData(todo.UserID)
}

func TestTodoShow(t *testing.T) {
	todo := CreateTodo()

	r := router.Router()

	req, err := http.NewRequest("GET", "/todos/"+strconv.Itoa(todo.ID), nil)
	if err != nil {
		log.Println(err.Error())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	DeleteData(todo.UserID)
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

	resData := UnmarshalData(w.Body)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, resData.Todo.Title, expectedTodo.Title)

	DeleteData(user.ID)
}
