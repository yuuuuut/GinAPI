package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type Tag entities.Tag
type TagIndexRes entities.TagIndexRes
type TagShowRes entities.TagShowRes

func (m Tag) GetAll() ([]TagIndexRes, error) {
	var (
		db   = util.GetDB()
		tags []Tag
	)

	if err := db.Where("parent_id = ?", "0").Preload("Tags").Find(&tags).Error; err != nil { //db.Not("id", 0).Find(&tags).Error; err != nil {
		return []TagIndexRes{}, err
	}

	var res []TagIndexRes
	q, err := json.Marshal(tags)
	if err != nil {
		return []TagIndexRes{}, err
	}

	if err := json.Unmarshal(q, &res); err != nil {
		return []TagIndexRes{}, err
	}

	return res, nil
}

func (m Tag) GetById(id string) (TagShowRes, error) {
	var (
		db  = util.GetDB()
		tag Tag
	)

	if err := db.Preload("Todos").Preload("Todos.User").First(&tag, id).Error; err != nil {
		return TagShowRes{}, err
	}

	var res TagShowRes
	q, err := json.Marshal(tag)
	if err != nil {
		return TagShowRes{}, err
	}

	if err := json.Unmarshal(q, &res); err != nil {
		return TagShowRes{}, err
	}

	return res, nil
}

func (m Tag) CreateM(c *gin.Context) error {
	var (
		db = util.GetDB()
	)

	tag1 := Tag{Name: "アウトドア", ParentID: nil}
	if err := db.Create(&tag1).Error; err != nil {
		return err
	}

	tag2 := Tag{Name: "キャンプ", ParentID: &tag1.ID}
	if err := db.Create(&tag2).Error; err != nil {
		return err
	}

	tag3 := Tag{Name: "インドア", ParentID: nil}
	if err := db.Create(&tag3).Error; err != nil {
		return err
	}

	tag4 := Tag{Name: "ゲーム", ParentID: &tag3.ID}
	if err := db.Create(&tag4).Error; err != nil {
		return err
	}

	return nil
}
