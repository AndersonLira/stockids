package model

//Log register each behaviour of child
type Log struct {
	ChildID string `json:"child_id"`
	Date    int64  `json:"date"`
	Message string `json:"message"`
	Score   int    `json:"score"`
}
