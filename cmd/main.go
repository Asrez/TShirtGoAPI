package main

import (
	"log"

	"github.com/Asrez/TShirtGoAPI/api"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/config"
)
func main(){
	cfg := config.GetConfig()
	err := db.InitDb(cfg)
	defer db.CloseDb()

	if err != nil {
		log.Fatal("Connection to database failed error :" , err)
	}

	api.InitServer()
}