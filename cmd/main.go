package main

import (
	"github.com/Asrez/TShirtGoAPI/api"
	"github.com/Asrez/TShirtGoAPI/config"
	"github.com/Asrez/TShirtGoAPI/data/cache"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/migration"
	"github.com/Asrez/TShirtGoAPI/pkg/logging"
)
func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := db.InitDb(cfg)
	defer db.CloseDb()

	if err != nil {
		logger.Fatal(logging.Postgres , logging.Startup , err.Error(),nil)
	}

	err = cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis , logging.Startup , err.Error(), nil)
	}
	migration.Up()
	api.InitServer()
}