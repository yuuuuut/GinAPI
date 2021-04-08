package models

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type Tag entities.Tag

func (m Tag) GetById(id string) (Tag, error) {
	var (
		db  = util.GetDB()
		tag Tag
	)

	if err := db.Preload("Todos").Preload("Todos.User").First(&tag, id).Error; err != nil {
		return tag, err
	}

	return tag, nil
}

func (m Tag) CreateM(c *gin.Context) error {
	var (
		db = util.GetDB()
	)

	tag1 := Tag{Name: "アウトドア"}
	if err := db.Create(&tag1).Error; err != nil {
		return err
	}

	tag2 := Tag{Name: "キャンプ", ParentID: tag1.ID}
	if err := db.Create(&tag2).Error; err != nil {
		return err
	}

	tag3 := Tag{Name: "インドア"}
	if err := db.Create(&tag3).Error; err != nil {
		return err
	}

	tag4 := Tag{Name: "ゲーム", ParentID: tag3.ID}
	if err := db.Create(&tag4).Error; err != nil {
		return err
	}

	return nil
}
