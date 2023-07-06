package handlers

import (
	"net/http"

	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/models"
	"github.com/gin-gonic/gin"
)


type Materials struct {

}

func NewMaterialsHandler() *Materials{
	return &Materials{}
}

func (c *Materials) Materials(ctx *gin.Context) {
	db := db.GetDb()

	var Materials []models.Materials
	result := db.Find(&Materials)
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
		"data":   Materials,
	})
}