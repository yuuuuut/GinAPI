package models

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type User entities.User

func (m User) GetById(id string) (User, error) {
	var db = util.GetDB()
	var user User

	if err := db.Where("id = ?", id).Preload("Todos").Preload("Todos.User").First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (m User) CreateM(c *gin.Context) (User, error) {
	var (
		db   = util.GetDB()
		user User
	)

	if err := c.BindJSON(&user); err != nil {
		return user, err
	}

	if err := db.Where(User{ID: user.ID}).Assign(User{DisplayName: user.DisplayName, PohotURL: user.PohotURL}).FirstOrCreate(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
