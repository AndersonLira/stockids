package main

//LogTotal represents votes agregation
type LogTotal struct {
	ChildID     string `json:"child_id"`
	Accumulated int    `json:"accumulated"`
}
