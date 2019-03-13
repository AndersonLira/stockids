package model

import (
	"time"
)

//Log register each behaviour of child
type Log struct {
	ID      string    `json:"id"`
	ChildID string    `json:"child_id"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
	Score   int       `json:"score"`
}
