package routes


import (
	"github.com/gin-gonic/gin"
	"github.com/Asrez/TShirtGoAPI/api/handlers"

)

func ProductsRoute(r *gin.RouterGroup){
	h := handlers.NewProductsHandler()
	r.GET("/",h.Products)
}