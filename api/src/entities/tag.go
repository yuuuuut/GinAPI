package entities

type Tag struct {
	ID       int
	Name     string
	ParentID int
	Todos    []Todo `gorm:"many2many:todos_tags"`
}
