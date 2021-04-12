package entities

type Todo struct {
	ID       int
	Title    string    `gorm:"not null" binding:"required"`
	Status   bool      `gorm:"default:false"`
	UserID   string    `gorm:"not null;REFERENCES users(id)"`
	User     User      `gorm:"PRELOAD:false"`
	Tags     []Tag     `gorm:"many2many:todos_tags"`
	Comments []Comment `gorm:"ForeignKey:TodoID"`
}

type CreateTodoReq struct {
	Title string
	Tags  []string
}
