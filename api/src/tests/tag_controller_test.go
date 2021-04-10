package tests

import (
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

type Tag entities.Tag

type IndexTag struct {
	Tags []Tag
}

type ShowTag struct {
	Tag struct {
		ID       int
		Name     string
		ParentID int
		Todos    []Todo
	}
}

func TestTagIndex(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo, todo.ID)
	defer DeleteData(user, user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/tags", nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	reqBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData IndexTag
	if err := json.Unmarshal(reqBody, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.Tags[0].ParentID, 0)
	assert.Equal(t, resData.Tags[0].Tags[0].ParentID, resData.Tags[0].ID)
}
func TestTagShow(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo, todo.ID)
	defer DeleteData(user, user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/tags/"+strconv.Itoa(todo.Tags[0].ID), nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	reqBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData ShowTag
	if err := json.Unmarshal(reqBody, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.Tag.Name, todo.Tags[0].Name)
	assert.Equal(t, resData.Tag.Todos[0].ID, todo.ID)
}
