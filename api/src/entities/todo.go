package entities

type Todo struct {
	ID     int
	Title  string `gorm:"not null" binding:"required"`
	Status bool   `gorm:"default:false"`
	UserID int    `gorm:"not null;REFERENCES users"`
	User   User
}
