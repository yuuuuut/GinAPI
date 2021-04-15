package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type User entities.User
type UserShowRes entities.UserShowRes
type UserCreateRes entities.UserCreateRes

func (m User) GetById(id, offset, limit string) (UserShowRes, error) {
	var (
		db   = util.GetDB()
		user User
	)

	if err := db.Where("id = ?", id).Preload("Todos", func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}).Preload("Todos.User").First(&user).Error; err != nil {
		return UserShowRes{}, err
	}

	var res UserShowRes
	data, err := json.Marshal(user)
	if err != nil {
		return UserShowRes{}, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return UserShowRes{}, err
	}

	return res, nil
}

func (m User) CreateM(c *gin.Context) (UserCreateRes, error) {
	var (
		db   = util.GetDB()
		user User
	)

	if err := c.BindJSON(&user); err != nil {
		return UserCreateRes{}, err
	}

	if err := db.Where(User{ID: user.ID}).Assign(User{DisplayName: user.DisplayName, PohotURL: user.PohotURL}).FirstOrCreate(&user).Error; err != nil {
		return UserCreateRes{}, err
	}

	var res UserCreateRes
	data, err := json.Marshal(user)
	if err != nil {
		return UserCreateRes{}, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return UserCreateRes{}, err
	}

	return res, nil
}
