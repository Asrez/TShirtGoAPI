package migration

import (
	"math/rand"
	"reflect"

	"github.com/Asrez/TShirtGoAPI/config"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/models"
	"github.com/Asrez/TShirtGoAPI/pkg/logging"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)


var logger = logging.NewLogger(config.GetConfig())


func Up(){
	database := db.GetDb()
	createTables(database)
	ExecuteDataFakerIfEmpty(database)
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
	tables = addNewTable(database , models.Products{} ,tables)


	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)

}

func CreateDataWithFaker(model interface{}, database *gorm.DB) {
	for i := 0; i < 10; i++ {
		record := reflect.New(reflect.TypeOf(model)).Elem()

		idField := record.FieldByName("ID")
		if idField.IsValid() && idField.CanSet() && idField.Kind() == reflect.Int {
			idField.SetInt(int64(i + 1))
		}

		nameField := record.FieldByName("Name")
		if nameField.IsValid() && nameField.CanSet() && nameField.Kind() == reflect.String {
			nameField.SetString(faker.Word())
		}

		database.Create(record.Addr().Interface())
	}
}

func CreateProductsDataWithFaker(database *gorm.DB) {
	for i := 0; i < 10; i++ {
		products := models.Products{
			BaseTable: models.BaseTable{
				Id:   i + 1,
				Name: faker.Word(),
			},
			CategoryID:  rand.Intn(10) + 1,
			BrandID:     rand.Intn(10) + 1,
			SizeID:      rand.Intn(10) + 1,
			ColorID:     rand.Intn(10) + 1,
			MaterialID:  rand.Intn(10) + 1,
			Price:       rand.Float64() * 100,
		}
		database.Create(&products)
	}
}


func ExecuteDataFakerIfEmpty(database *gorm.DB) {
	var count int64

	database.Model(&models.Categories{}).Count(&count)
	if count == 0 {
		CreateDataWithFaker(models.Categories{}, database)
	}

	database.Model(&models.Brands{}).Count(&count)
	if count == 0 {
		CreateDataWithFaker(models.Brands{}, database)
	}

	database.Model(&models.Sizes{}).Count(&count)
	if count == 0 {
		CreateDataWithFaker(models.Sizes{}, database)
	}

	database.Model(&models.Colors{}).Count(&count)
	if count == 0 {
		CreateDataWithFaker(models.Colors{}, database)
	}

	database.Model(&models.Materials{}).Count(&count)
	if count == 0 {
		CreateDataWithFaker(models.Materials{}, database)
	}

	database.Model(&models.Products{}).Count(&count)
	if count == 0 {
		CreateProductsDataWithFaker(database)
	}
}
