package main

import (
	"time"
)

//Log register each behaviour of child
type Log struct {
	ChildID string    `json:"child_id"`
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
	Partial int       `json:"partial"`
	Score   int       `json:"score"`
}
