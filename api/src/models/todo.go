package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type Todo entities.Todo
type TodoIndexRes entities.TodoIndexRes
type TodoShowRes entities.TodoShowRes
type TodoCreateRes entities.TodoCreateRes
type TodoUpdateRes entities.TodoUpdateRes
type TodoDeleteRes entities.TodoDeleteRes

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

	if err := json.Unmarshal(data, &res); err != nil {
		return []TodoIndexRes{}, err
	}

	return res, nil
}

func (m Todo) GetById(id string) (TodoShowRes, error) {
	var (
		db   = util.GetDB()
		todo Todo
	)

	if err := db.First(&todo, id).Related(&todo.User).Related(&todo.Tags, "Tags").Error; err != nil {
		return TodoShowRes{}, err
	}

	var res TodoShowRes

	data, err := json.Marshal(todo)
	if err != nil {
		return TodoShowRes{}, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return TodoShowRes{}, err
	}

	return res, nil
}

func (m Todo) CreateM(c *gin.Context) (TodoCreateRes, map[string]string, error) {
	var (
		db     = util.GetDB()
		userId = c.GetString("currentUserId")
		user   User
		req    CreateTodoReq
		tags   []entities.Tag
	)

	if err := c.BindJSON(&req); err != nil {
		errorMessages := util.TodoValidation(err)
		return TodoCreateRes{}, errorMessages, err
	}

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return TodoCreateRes{}, nil, err
	}

	if err := db.Where("id IN (?)", req.Tags).Find(&tags).Error; err != nil {
		return TodoCreateRes{}, nil, err
	}

	todo := Todo{Title: req.Title, UserID: userId, User: entities.User(user), Tags: tags}
	if err := db.Create(&todo).Error; err != nil {
		return TodoCreateRes{}, nil, err
	}

	var res TodoCreateRes

	data, err := json.Marshal(todo)
	if err != nil {
		return TodoCreateRes{}, nil, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return TodoCreateRes{}, nil, err
	}

	return res, nil, nil
}

func (m Todo) UpdateById(id string, c *gin.Context) (TodoUpdateRes, error) {
	var (
		db   = util.GetDB()
		todo Todo
	)

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return TodoUpdateRes{}, err
	}

	todo.Status = true

	if err := db.Save(&todo).Error; err != nil {
		return TodoUpdateRes{}, err
	}

	var res TodoUpdateRes

	data, err := json.Marshal(todo)
	if err != nil {
		return TodoUpdateRes{}, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return TodoUpdateRes{}, err
	}

	return res, nil
}

func (m Todo) DeleteById(id string) (TodoDeleteRes, error) {
	var (
		db   = util.GetDB()
		todo Todo
	)

	if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return TodoDeleteRes{}, err
	}

	var res TodoDeleteRes

	data, err := json.Marshal(todo)
	if err != nil {
		return TodoDeleteRes{}, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return TodoDeleteRes{}, err
	}

	return res, nil
}
