package main

//ChildTotal represents votes agregation
type ChildTotal struct {
	ChildID     string `json:"child_id"`
	Accumulated int    `json:"accumulated"`
}
