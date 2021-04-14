package entities

type Todo struct {
	ID       int       `json:"id"`
	Title    string    `gorm:"not null" binding:"required" json:"title"`
	Status   bool      `gorm:"default:false" json:"status"`
	UserID   string    `gorm:"not null;REFERENCES users(id)" json:"user_id"`
	User     User      `gorm:"PRELOAD:false" json:"user"`
	Tags     []Tag     `gorm:"many2many:todos_tags" json:"tags"`
	Comments []Comment `gorm:"ForeignKey:TodoID" json:"comments"`
}

type TodoIndexRes struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	UserID string `json:"user_id"`
	User   struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
	} `json:"user"`
}

type TodoShowRes struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	UserID string `json:"user_id"`
	User   struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
	} `json:"user"`
	Tags []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		ParentID int    `json:"parent_id"`
	} `json:"tags"`
}

type CreateTodoReq struct {
	Title string
	Tags  []string
}
