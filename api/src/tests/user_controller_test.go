package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/router"
)

type User entities.User

type UserShow struct {
	User struct {
		ID          string
		DisplayName string
		PohotURL    string
		Todos       []Todo
	}
}

type UserCreate struct {
	User struct {
		ID          string
		DisplayName string
		PohotURL    string
		Todos       []Todo
	}
	Profile struct {
		ID       int
		Nickname string
		Sex      string
		Age      int
		UserID   string
	}
}

func TestUserShow(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo, todo.ID)
	defer DeleteData(user, user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/users/"+user.ID, nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	reqBody, _ := ioutil.ReadAll(w.Body)

	var resData UserShow
	if err := json.Unmarshal(reqBody, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.User.Todos[0].ID, todo.ID)
	assert.Equal(t, resData.User.Todos[0].User, user)
}

func TestUserCreate(t *testing.T) {
	uid := os.Getenv("FIREBASE_ADMIN_USER_UID")
	defer DeleteData(entities.User{}, uid)

	user := new(User)
	user.ID = uid
	user.DisplayName = "TestUser"
	user.PohotURL = "TestPhoto"
	user_json, _ := json.Marshal(user)
	body := bytes.NewBuffer(user_json)

	r := router.Router()
	req, err := http.NewRequest("POST", "/users", body)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	b, _ := ioutil.ReadAll(w.Body)

	var resData UserCreate
	if err := json.Unmarshal(b, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, resData.User.ID, uid)
	assert.Equal(t, resData.Profile.UserID, uid)
}
