package models

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type Todo entities.Todo

func (m Todo) GetAll(offset, limit string) ([]Todo, error) {
	var db = util.GetDB()
	var todos []Todo

	if err := db.Offset(offset).Limit(limit).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (m Todo) GetById(id string) (Todo, error) {
	var db = util.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (m Todo) CreateM(c *gin.Context) (Todo, map[string]string, error) {
	var (
		db   = util.GetDB()
		user User
		req  Todo
	)

	userId := c.GetString("currentUserId")

	if err := c.BindJSON(&req); err != nil {
		errorMessages := util.TodoValidation(err)
		return req, errorMessages, err
	}

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return req, nil, err
	}

	todo := Todo{Title: req.Title, UserID: userId, User: entities.User(user)}
	if err := db.Create(&todo).Error; err != nil {
		return todo, nil, err
	}

	return todo, nil, nil
}

func (m Todo) UpdateById(id string, c *gin.Context) (Todo, error) {
	var db = util.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	}

	todo.Status = true

	if err := db.Save(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (m Todo) DeleteById(id string) (Todo, error) {
	var db = util.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}
