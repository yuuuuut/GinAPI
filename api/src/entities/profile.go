package entities

type Profile struct {
	ID       int `gorm:"primary_key"`
	Nickname string
	Sex      string
	Age      int
	UserID   string `gorm:"unique;not null"`
}
