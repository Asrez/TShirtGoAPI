package routes


import (
	"github.com/gin-gonic/gin"
	"github.com/Asrez/TShirtGoAPI/api/handlers"

)

func MaterialsRoute(r *gin.RouterGroup){
	h := handlers.NewMaterialsHandler()
	r.GET("/",h.Materials)
}