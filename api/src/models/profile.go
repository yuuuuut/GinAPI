package models

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type Profile entities.Profile

func (m Profile) GetById(id string) (Profile, error) {
	var db = util.GetDB()
	var profile Profile

	if err := db.Where("user_id = ?", id).First(&profile).Error; err != nil {
		return profile, err
	}

	return profile, nil
}

func (m Profile) CreateM(c *gin.Context, userId string) (Profile, error) {
	var (
		db      = util.GetDB()
		profile Profile
	)

	if err := db.Where(Profile{UserID: userId}).FirstOrCreate(&profile).Error; err != nil {
		return profile, err
	}

	return profile, nil
}
