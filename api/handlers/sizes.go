package handlers

import (
	"net/http"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/models"
	"github.com/gin-gonic/gin"
)


type Sizes struct {

}

func NewSizesHandler() *Sizes{
	return &Sizes{}
}

func (c *Sizes) Sizes(ctx *gin.Context) {
	db := db.GetDb()

	var Sizes []models.Sizes
	result := db.Find(&Sizes)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
		"data":   Sizes,
	})
}