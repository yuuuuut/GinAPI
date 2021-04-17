package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type Comment entities.Comment
type CommentCreateReq entities.CommentCreateReq
type CommentCreateRes entities.CommentCreateRes

func (m Comment) CreateM(c *gin.Context) (CommentCreateRes, error) {
	var (
		db      = util.GetDB()
		userId  = c.GetString("currentUserId")
		req     CommentCreateReq
		user    User
		todo    Todo
		comment Comment
	)

	if err := c.BindJSON(&req); err != nil {
		return CommentCreateRes{}, err
	}

	if err := db.Where("id = ?", userId).Select("id").First(&user).Error; err != nil {
		return CommentCreateRes{}, err
	}

	if err := db.Where("id = ?", req.TodoID).Find(&todo).Error; err != nil {
		return CommentCreateRes{}, err
	}

	if req.ParentID == 0 {
		comment = Comment{Content: req.Content, ParentID: nil, UserID: userId, TodoID: req.TodoID, User: entities.User(user)}
		if err := db.Create(&comment).Error; err != nil {
			return CommentCreateRes{}, err
		}
	} else {
		comment = Comment{Content: req.Content, ParentID: &req.ParentID, UserID: userId, TodoID: req.TodoID, User: entities.User(user)}
		if err := db.Create(&comment).Error; err != nil {
			return CommentCreateRes{}, err
		}
	}

	var res CommentCreateRes
	q, err := json.Marshal(comment)
	if err != nil {
		return CommentCreateRes{}, err
	}

	if err := json.Unmarshal(q, &res); err != nil {
		return CommentCreateRes{}, err
	}

	return res, nil
}
