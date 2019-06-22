package dbsetup

import (
	"log"

	"github.com/andersonlira/godyn/db"

	"github.com/andersonlira/stockids/model"
)

var tables = []model.Tableable{
	&model.Child{},
}

//CreateTables of application
func CreateTables() {

	for _, t := range tables {
		_, err := db.CreateTable(t)
		if err != nil {
			log.Println(err)
		}
	}
}
