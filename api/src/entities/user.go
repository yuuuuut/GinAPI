package entities

type User struct {
	ID          string  `gorm:"primary_key" json:"id"`
	DisplayName string  `gorm:"not null" json:"display_name"`
	PohotURL    string  `gorm:"not null" json:"photo_url"`
	Todos       []Todo  `gorm:"ForeignKey:UserID" json:"todos"`
	Profile     Profile `gorm:"ForeignKey:UserID" json:"profile"`
}

type UserShowRes struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	PohotURL    string `json:"photo_url"`
	Todos       []struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Status bool   `json:"status"`
		UserID string `json:"user_id"`
		User   struct {
			ID          string `json:"id"`
			DisplayName string `json:"display_name"`
		} `json:"user"`
	} `json:"todos"`
}

type UserCreateRes struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	PohotURL    string `json:"photo_url"`
}
