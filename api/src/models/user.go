package models

import (
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
