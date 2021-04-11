package models

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/entities"
	"github.com/yuuuuut/gin-api/src/util"
)

type Comment entities.Comment
type CreateCommentReq entities.CreateCommentReq

func (m Comment) CreateM(c *gin.Context) (Comment, error) {
	var (
		db      = util.GetDB()
		userId  = c.GetString("currentUserId")
		req     CreateCommentReq
		user    User
		todo    Todo
		comment Comment
	)

	if err := c.BindJSON(&req); err != nil {
		return Comment{}, err
	}

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return Comment{}, err
	}

	if err := db.Where("id = ?", req.TodoID).Find(&todo).Error; err != nil {
		return Comment{}, err
	}

	if req.ParentID == 0 {
		comment = Comment{Content: req.Content, ParentID: nil, UserID: userId, TodoID: req.TodoID, User: entities.User(user), Todo: entities.Todo(todo)}
		if err := db.Create(&comment).Error; err != nil {
			return comment, err
		}
	} else {
		comment = Comment{Content: req.Content, ParentID: &req.ParentID, UserID: userId, TodoID: req.TodoID, User: entities.User(user), Todo: entities.Todo(todo)}
		if err := db.Create(&comment).Error; err != nil {
			return comment, err
		}
	}

	return comment, nil
}
