package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/router"
)

type Profile entities.Profile

type ShowProfile struct {
	Profile struct {
		ID       int
		Nickname string
		Sex      string
		Age      int
		UserID   string
	}
}

func TestProfileShow(t *testing.T) {
	user := CreateUser()
	profile := CreateProfile(user.ID)
	defer DeleteData(profile, profile.ID)
	defer DeleteData(user, user.ID)

	r := router.Router()
	req, err := http.NewRequest("GET", "/profiles/"+user.ID, nil)
	if err != nil {
		log.Println(err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	reqBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resData ShowProfile
	if err := json.Unmarshal(reqBody, &resData); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, resData.Profile.UserID, user.ID)
}
