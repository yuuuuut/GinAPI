package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type Todo entities.Todo
type TodoIndexRes entities.TodoIndexRes
type CreateTodoReq entities.CreateTodoReq

func (m Todo) GetAll(offset, limit string) ([]TodoIndexRes, error) {
	var (
		db    = util.GetDB()
		todos []Todo
	)

	if err := db.Offset(offset).Limit(limit).Preload("User").Find(&todos).Error; err != nil {
		return []TodoIndexRes{}, err
	}

	var res []TodoIndexRes

	data, err := json.Marshal(todos)
	if err != nil {
		return []TodoIndexRes{}, err
	}

	if err := json.Unmarshal([]byte(data), &res); err != nil {
		return []TodoIndexRes{}, err
	}

	return res, nil
}

func (m Todo) GetById(id string) (Todo, error) {
	var db = util.GetDB()
	var todo Todo

	if err := db.First(&todo, id).Related(&todo.User).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (m Todo) CreateM(c *gin.Context) (Todo, map[string]string, error) {
	var (
		db   = util.GetDB()
		user User
		req  CreateTodoReq
		tags []entities.Tag
	)

	userId := c.GetString("currentUserId")

	if err := c.BindJSON(&req); err != nil {
		errorMessages := util.TodoValidation(err)
		return Todo{}, errorMessages, err
	}

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return Todo{}, nil, err
	}

	if err := db.Where("id IN (?)", req.Tags).Find(&tags).Error; err != nil {
		return Todo{}, nil, err
	}

	todo := Todo{Title: req.Title, UserID: userId, User: entities.User(user), Tags: tags}
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
