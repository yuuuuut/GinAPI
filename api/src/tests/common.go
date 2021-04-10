package tests

import (
	"log"
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

func CreateProfile(userId string) entities.Profile {
	var db = util.GetDB()

	profile := entities.Profile{UserID: userId}
	if err := db.Create(&profile).Error; err != nil {
		panic(err.Error())
	}

	return profile
}

func CreateTodo(userId string) entities.Todo {
	var (
		db   = util.GetDB()
		user entities.User
		tags []entities.Tag
	)

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Where("id IN (?)", []string{"1", "2"}).Find(&tags).Error; err != nil {
		log.Fatal(err.Error())
	}

	todo := entities.Todo{Title: "TestTodo", UserID: userId, User: entities.User(user), Tags: tags}
	if err := db.Create(&todo).Error; err != nil {
		panic(err.Error())
	}

	return todo
}

func DeleteData(t interface{}, id interface{}) {
	var db = util.GetDB()

	switch t.(type) {
	case entities.Todo:
		if err := db.Where("id = ?", id).Delete(entities.Todo{}).Error; err != nil {
			panic(err.Error())
		}
	case entities.User:
		if err := db.Where("id = ?", id).Delete(entities.User{}).Error; err != nil {
			panic(err.Error())
		}
	case entities.Profile:
		if err := db.Where("id = ?", id).Delete(entities.Profile{}).Error; err != nil {
			panic(err.Error())
		}
	}
}
