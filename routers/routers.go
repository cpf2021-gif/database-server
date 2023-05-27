package routers

import (
	"github.com/gin-gonic/gin"

	"server/app/user"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 中间件

	// 路由
	// 注册用户路由
	user.Routers(r)

	return r
}
