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
		
		brands := v1.Group("/brands")
		routes.BrandsRoute(brands)

		sizes := v1.Group("/sizes")
		routes.BrandsRoute(sizes)

		colors := v1.Group("/colors")
		routes.ColorsRoute(colors)

		materials := v1.Group("/materials")
		routes.BrandsRoute(materials)

		products := v1.Group("/products")
		routes.ProductsRoute(products)
	}

	r.Run(fmt.Sprintf(":%s", config.Server.Port))
}