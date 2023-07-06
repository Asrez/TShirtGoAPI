package routes


import (
	"github.com/gin-gonic/gin"
	"github.com/Asrez/TShirtGoAPI/api/handlers"

)

func BrandsRoute(r *gin.RouterGroup){
	h := handlers.NewBrandsHandler()
	r.GET("/",h.Brands)
}