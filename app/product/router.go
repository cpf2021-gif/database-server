package product

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/products", GetProducts)
	e.GET("/suppliers", GetSuppliers)
	e.GET("/sellers", GetSellers)

	e.POST("/products", CreateProduct)
	e.POST("/suppliers", CreateSupplier)
	e.POST("/sellers", CreateSeller)

	e.PATCH("/products/:id", UpdateProduct)
	e.PATCH("/suppliers/:name", UpdateSupplier)
	e.PATCH("/sellers/:name", UpdateSeller)
}
