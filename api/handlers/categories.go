package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type Categories struct {

}

func NewCategoriesHandler() *Categories{
	return &Categories{}
}

func (u *Categories) Categories(c *gin.Context) {
	c.JSON(http.StatusOK,"Categories ... ")
	return
}