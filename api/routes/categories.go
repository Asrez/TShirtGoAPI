package routes


import (
	"github.com/gin-gonic/gin"
	"github.com/Asrez/TShirtGoAPI/api/handlers"

)

func CategoriesRoute(r *gin.RouterGroup){
	h := handlers.NewCategoriesHandler()
	r.GET("/",h.Categories)
}