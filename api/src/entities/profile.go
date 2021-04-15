package entities

type Profile struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	UserID   string `gorm:"unique;not null" json:"user_id"`
}
