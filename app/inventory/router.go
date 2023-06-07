package inventory

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/inventories", GetInventories)
	e.GET("/inbounds", GetInBounds)
	e.GET("/outbounds", GetOutBounds)

	e.PUT("/inbounds/export", ExportInbound)
	e.PUT("/outbounds/export", ExportOutbound)

	e.POST("/inbounds", CreateInbound)
	e.POST("/outbounds", CreateOutbound)
	e.PATCH("/inventories/:id", UpdateInventoryRequestByID)
}
