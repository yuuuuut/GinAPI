package todo

import (
	"github.com/yuuuuut/gin-api/src/db"
	"github.com/yuuuuut/gin-api/src/entity"
)

type Model struct{}

type Todo entity.Todo

func (m Model) GetAll() ([]Todo, error) {
	db := db.GetDB()

	var todos []Todo

	err := db.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}
