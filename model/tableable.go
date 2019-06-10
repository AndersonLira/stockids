package model

//Tableable is the interface for structs those should be a model
type Tableable interface {
	GetTableName() string
}
