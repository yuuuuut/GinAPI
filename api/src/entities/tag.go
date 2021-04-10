package entities

type Tag struct {
	ID       int
	Name     string
	ParentID int
	Tags     []Tag  `gorm:"ForeignKey:ParentID"`
	Todos    []Todo `gorm:"many2many:todos_tags"`
}
