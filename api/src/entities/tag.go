package entities

type Tag struct {
	ID       int
	Name     string
	ParentID *int   `json:"parent_id"`
	Tags     []Tag  `gorm:"ForeignKey:ParentID"`
	Todos    []Todo `gorm:"many2many:todos_tags"`
}

type TagIndexRes struct {
	ID       int
	Name     string
	ParentID int
	Tags     []struct {
		ID       int
		Name     string
		ParentID int
	}
}
