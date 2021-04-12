package entities

type User struct {
	ID          string  `gorm:"primary_key"`
	DisplayName string  `gorm:"not null"`
	PohotURL    string  `gorm:"not null"`
	Todos       []Todo  `gorm:"ForeignKey:UserID"`
	Profile     Profile `gorm:"ForeignKey:UserID"`
}
