package main

import (
	"log"

	"github.com/Asrez/TShirtGoAPI/api"
	"github.com/Asrez/TShirtGoAPI/data/db"
)
func main(){
	
	err := db.InitDb()
	defer db.CloseDb()

	if err != nil {
		log.Fatal("Connection to database failed error :" , err)
	}

	api.InitServer()
}