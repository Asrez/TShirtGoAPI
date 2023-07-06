package handlers

import (
	"net/http"

	helper "github.com/Asrez/TShirtGoAPI/api/helpers"
	"github.com/Asrez/TShirtGoAPI/data/db"
	"github.com/Asrez/TShirtGoAPI/data/models"
	"github.com/gin-gonic/gin"
)


type Brands struct {

}

func NewBrandsHandler() *Brands{
	return &Brands{}
}

func (c *Brands) Brands(ctx *gin.Context) {
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

	ctx.JSON(http.StatusOK,helper.GenerateBaseResponse(gin.H{
		"result": "success",
		"data":   brands,
	} , true ,200))
}