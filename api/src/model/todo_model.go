package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/db"
	"github.com/yuuuuut/gin-api/src/entity"
	"github.com/yuuuuut/gin-api/src/util"
)

type Model struct{}

type Todo entity.Todo

func (m Model) GetAll() ([]Todo, error) {
	db := db.GetDB()
	var todos []Todo

	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (m Model) Get(id string) (Todo, error) {
	db := db.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (m Model) CreateM(c *gin.Context) (Todo, error, map[string]string) {
	db := db.GetDB()
	var todo Todo

	if err := c.BindJSON(&todo); err != nil {
		errorMessages := util.TodoValidation(err)
		return todo, err, errorMessages
	}

	if err := db.Create(&todo).Error; err != nil {
		return todo, err, nil
	}

	return todo, nil, nil
}
