package entities

type Comment struct {
	ID       int    `gorm:"primary_key"`
	Content  string `gorm:"not null"`
	ParentID int
	UserID   string `gorm:"not null"`
	TodoID   int    `gorm:"not null"`
	User     User
	Todo     Todo
	Comments []Comment `gorm:"ForeignKey:ParentID"`
}
