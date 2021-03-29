package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"strings"
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/router"
	"github.com/yuuuuut/gin-api/src/util"
)

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
	r := router.Router()

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		log.Println(err.Error())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	fmt.Println(w)

	assert.Equal(t, 200, w.Code)
}

func TestTodoPost(t *testing.T) {
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

	fmt.Println(resData)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, resData.Todo.Title, expectedTodo.Title)
}
