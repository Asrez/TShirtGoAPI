package routes


import (
	"github.com/gin-gonic/gin"
	"github.com/Asrez/TShirtGoAPI/api/handlers"

)

func ColorsRoute(r *gin.RouterGroup){
	h := handlers.NewColorsHandler()
	r.GET("/",h.Colors)
}