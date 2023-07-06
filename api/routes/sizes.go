package routes


import (
	"github.com/gin-gonic/gin"
	"github.com/Asrez/TShirtGoAPI/api/handlers"

)

func SizesRoute(r *gin.RouterGroup){
	h := handlers.NewSizesHandler()
	r.GET("/",h.Sizes)
}