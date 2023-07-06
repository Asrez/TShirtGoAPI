package handlers

import (
	"net/http"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/models"
	"github.com/gin-gonic/gin"
)


type Colors struct {

}

func NewColorsHandler() *Colors{
	return &Colors{}
}

func (c *Colors) Colors(ctx *gin.Context) {
	db := db.GetDb()

	var Colors []models.Colors
	result := db.Find(&Colors)
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
		"data":   Colors,
	})
}