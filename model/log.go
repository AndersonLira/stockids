package model

//Family represents one group
type Family struct {
	ChildID string `json:"child_id"`
	Date    int64  `json:"date"`
	Message string `json:"message"`
	Score   int    `json:"score"`
}
