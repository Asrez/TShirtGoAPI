package migration

import (
	"github.com/Asrez/TShirtGoAPI/config"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/models"
	"github.com/Asrez/TShirtGoAPI/pkg/logging"
	"gorm.io/gorm"
)


var logger = logging.NewLogger(config.GetConfig())


func Up(){
	database := db.GetDb()
	createTables(database)
	logger.Info(logging.Postgres, logging.Migration, "UP", nil)

}


func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createTables(database *gorm.DB){
	tables := []interface{}{}

	tables = addNewTable(database , models.Categories{} , tables)
	tables = addNewTable(database , models.Brands{} , tables)
	tables = addNewTable(database , models.Sizes{} , tables)
	tables = addNewTable(database , models.Colors{} , tables)
	tables = addNewTable(database , models.Materials{} , tables)




	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)

}