package entities

type User struct {
	ID          int
	DisplayName string `gorm:"not null"`
	PohotURL    string `gorm:"not null"`
	Todos       []Todo `gorm:"ForeignKey:UserID"`
}
