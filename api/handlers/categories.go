package handlers

import (
	"net/http"

	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/models"
	"github.com/gin-gonic/gin"
)


type Categories struct {

}

func NewCategoriesHandler() *Categories{
	return &Categories{}
}

func (c *Categories) Categories(ctx *gin.Context) {
	db := db.GetDb()

	var brands []models.Brands
	result := db.Find(&brands)
	if result.Error != nil {
		// Handle the error if any
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
		"data":   brands,
	})
}