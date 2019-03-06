package model

//Child is a model
type Child struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
	Name     string `json:"name"`
	MaxScore int    `json:"max_score"`
}
