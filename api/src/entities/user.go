package entities

type User struct {
	UID         string `gorm:"primary_key"`
	DisplayName string `gorm:"not null"`
	PohotURL    string `gorm:"not null"`
	Todos       []Todo `gorm:"ForeignKey:UserID"`
}
