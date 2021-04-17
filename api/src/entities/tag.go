package entities

type Tag struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
	Tags     []Tag  `gorm:"ForeignKey:ParentID" json:"tags"`
	Todos    []Todo `gorm:"many2many:todos_tags" json:"todos"`
}

type TagIndexRes struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
	Tags     []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		ParentID int    `json:"parent_id"`
	} `json:"tags"`
}

type TagShowRes struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
	Todos    []struct {
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
