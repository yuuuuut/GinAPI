package tests

import (
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
	"github.com/yuuuuut/gin-api/src/router"
)

type OneUser struct {
	User struct {
		ID          string
		DisplayName string
		PohotURL    string
		Todos       []Todo
	}
}

func TestUserShow(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo.ID)
	defer DeleteData(user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/users/"+user.ID, nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	reqBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData OneUser
	if err := json.Unmarshal(reqBody, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.User.Todos[0].ID, todo.ID)
	assert.Equal(t, resData.User.Todos[0].User, user)
}

func TestUserCreate(t *testing.T) {
	r := router.Router()
	uid := os.Getenv("FIREBASE_ADMIN_USER_UID")

	defer DeleteData(uid)

	reqBody := strings.NewReader(fmt.Sprintf(`{"ID": "%s","DisplayName":"TestUser","PohotURL":"TestPhoto"}`, uid))
	req, err := http.NewRequest("POST", "/users", reqBody)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	b, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData OneUser
	if err := json.Unmarshal(b, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, resData.User.ID, uid)
}
