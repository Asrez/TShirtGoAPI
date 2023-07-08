package handlers

import (
	"net/http"

	"github.com/Asrez/TShirtGoAPI/data/cache"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/models"
	"github.com/gin-gonic/gin"
)


type Products struct {

}

func NewProductsHandler() *Products{
	return &Products{}
}

func (c *Products) Products(ctx *gin.Context) {
	db := db.GetDb()
	redis := cache.GetRedis()

	values , err := cache.Get(redis , "products")

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "success",
			"data":   values,
		})
	}

	var Products []models.Products
	err = db.Preload("Category").Preload("Brand").Preload("Size").Preload("Color").Preload("Material").Find(&Products).Error
	if err != nil {
		cache.Set(redis,"products",err,10000)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
			"message": err,
		})
		return
	}
}