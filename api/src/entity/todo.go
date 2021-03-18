package entity

type Todo struct {
	ID     int
	Title  string `gorm:"not null" binding:"required"`
	Status bool   `gorm:"default:false"`
	/*
		UserID int `gorm:"TYPE:integer REFERENCES users"`
		User   User
	*/
}
