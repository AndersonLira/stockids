package model

//Child is a model
type Child struct {
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}
