package inventory

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/inventories", GetInventories)
	e.GET("/inbounds", GetInBounds)
	e.GET("/outbounds", GetOutBounds)

	e.POST("/inbounds", CreateInbound)
	e.POST("/outbounds", CreateOutbound)
}
