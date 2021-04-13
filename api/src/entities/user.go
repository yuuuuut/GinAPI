package entities

type User struct {
	ID          string  `gorm:"primary_key" json:"id"`
	DisplayName string  `gorm:"not null" json:"display_name"`
	PohotURL    string  `gorm:"not null" json:"photo_url"`
	Todos       []Todo  `gorm:"ForeignKey:UserID" json:"todos"`
	Profile     Profile `gorm:"ForeignKey:UserID" json:"profile"`
}
