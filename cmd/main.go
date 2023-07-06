package main

import (
	"github.com/Asrez/TShirtGoAPI/api"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/config"
	"github.com/Asrez/TShirtGoAPI/pkg/logging"
	"github.com/Asrez/TShirtGoAPI/db/migration"
)
func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := db.InitDb(cfg)
	defer db.CloseDb()

	if err != nil {
		// log.Fatal("Connection to database failed error :" , err)
		logger.Fatal(logging.Postgres , logging.Startup , err.Error() , nil)
	}
	migration.Up()
	api.InitServer()
}