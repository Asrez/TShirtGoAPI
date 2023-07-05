package api

import (
	"fmt"

	"github.com/Asrez/TShirtGoAPI/config"
	"github.com/gin-gonic/gin"
)

func InitServer(){
	config := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger()  ,gin.Recovery())



	r.Run(fmt.Sprintf(":%s", config.Server.Port))
}