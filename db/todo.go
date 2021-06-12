package db

type Todo struct {
	ID    string `json:"id" gorm:"type:uuid;primaryKey;"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
