package tests

import (
	"os"

	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

func CreateUser() entities.User {
	var db = util.GetDB()

	uid := os.Getenv("FIREBASE_ADMIN_USER_UID")

	user := entities.User{ID: uid, DisplayName: "TestName", PohotURL: "TestPhotoURL"}
	if err := db.Create(&user).Error; err != nil {
		panic(err.Error())
	}

	return user
}

func CreateTodo(userId string) entities.Todo {
	var db = util.GetDB()

	todo := entities.Todo{Title: "TestTodo", UserID: userId}
	if err := db.Create(&todo).Error; err != nil {
		panic(err.Error())
	}

	return todo
}

func DeleteData(id interface{}) {
	var db = util.GetDB()

	switch id.(type) {
	case int:
		if err := db.Where("id = ?", id).Delete(entities.Todo{}).Error; err != nil {
			panic(err.Error())
		}
	case string:
		if err := db.Where("id = ?", id).Delete(entities.User{}).Error; err != nil {
			panic(err.Error())
		}
	}
}
