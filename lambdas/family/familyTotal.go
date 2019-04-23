package main

//FamilyTotal represents votes agregation
type FamilyTotal struct {
	ChildID     string `json:"child_id"`
	Accumulated int    `json:"accumulated"`
}
