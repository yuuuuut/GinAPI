package entities

type Comment struct {
	ID       int    `gorm:"primary_key"`
	Content  string `gorm:"not null"`
	ParentID *int
	UserID   string `gorm:"not null"`
	TodoID   int    `gorm:"not null"`
	User     User
	Todo     Todo      `gorm:"PRELOAD:false"`
	Comments []Comment `gorm:"ForeignKey:ParentID"`
}

type CreateCommentReq struct {
	Content  string
	ParentID int
	TodoID   int
}

type CommentCreateRes struct {
	ID       int
	Content  string
	ParentID int
	UserID   string
	TodoID   int
	User     struct {
		ID          string
		DisplayName string
	}
}
