package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entity"
	"github.com/yuuuuut/gin-api/src/util"
)

type Model struct{}

type Todo entity.Todo

func (m Model) GetAll() ([]Todo, error) {
	db := util.GetDB()
	var todos []Todo

	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (m Model) GetById(id string) (Todo, error) {
	db := util.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (m Model) CreateM(c *gin.Context) (Todo, map[string]string, error) {
	db := util.GetDB()
	var todo Todo

	if err := c.BindJSON(&todo); err != nil {
		errorMessages := util.TodoValidation(err)
		return todo, errorMessages, err
	}

	if err := db.Create(&todo).Error; err != nil {
		return todo, nil, err
	}

	return todo, nil, nil
}

func (m Model) UpdateById(id string, c *gin.Context) (Todo, error) {
	db := util.GetDB()
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

func (m Model) DeleteById(id string) (Todo, error) {
	db := util.GetDB()
	var todo Todo

	if err := db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}
