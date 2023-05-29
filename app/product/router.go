package product

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.POST("/products", CreateProduct)
	e.GET("/products", GetProducts)
	e.GET("/suppliers", GetSuppliers)
	e.PATCH("/products/:id", UpdateProduct)
	e.POST("/suppliers", CreateSupplier)
}
