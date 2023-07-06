package api

import (
	"fmt"
	"github.com/Asrez/TShirtGoAPI/api/routes"
	"github.com/Asrez/TShirtGoAPI/config"
	"github.com/gin-gonic/gin"
)

func InitServer(){
	config := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger()  ,gin.Recovery())

	v1 := r.Group("api/v1")
	{
		categories := v1.Group("/categories")
		routes.CategoriesRoute(categories)
	}

	r.Run(fmt.Sprintf(":%s", config.Server.Port))
}