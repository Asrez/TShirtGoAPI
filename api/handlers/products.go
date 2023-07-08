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
	var Products []models.Products
	value , err := cache.GetString(redis,"products")

	if err != nil {

		ctx.JSON(http.StatusOK, gin.H{
			"result": "success with redis",
			"data":   value,
		})
	}

	err = db.Preload("Category").Preload("Brand").Preload("Size").Preload("Color").Preload("Material").Find(&Products).Error
	cache.SetString(redis,"products", Products , 1000)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": "success without redis",
		"data":   Products,
	})

	return
}