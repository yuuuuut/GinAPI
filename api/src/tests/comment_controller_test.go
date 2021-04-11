package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/router"
)

type Comment entities.Comment

type CreateComment struct {
	Comment struct {
		ID       int
		Content  string
		ParentID *int
		UserID   string
		TodoID   int
		User     User
	}
}

func TestCommentPost(t *testing.T) {
	user := CreateUser()
	todo := CreateTodo(user.ID)
	defer DeleteData(todo, todo.ID)
	defer DeleteData(user, user.ID)

	comment := new(Comment)
	comment.Content = "TestContent"
	comment.ParentID = nil
	comment.TodoID = todo.ID
	comment_json, _ := json.Marshal(comment)
	body := bytes.NewBuffer(comment_json)

	r := router.Router()
	req, err := http.NewRequest("POST", "/comments", body)
	if err != nil {
		fmt.Println(err.Error())
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	b, _ := ioutil.ReadAll(w.Body)

	var resData CreateComment
	if err := json.Unmarshal(b, &resData); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resData.Comment)

	assert.Equal(t, 201, w.Code)
}
