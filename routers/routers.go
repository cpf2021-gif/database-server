package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/app/inventory"
	"server/app/product"
	"server/app/user"
	"server/global"
	"server/util"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 中间件

	// 路由
	// 备份数据库
	r.GET("/backup", func(c *gin.Context) {
		err := util.Backup(global.GL_VIPER)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "备份失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "备份成功",
		})
	})

	// 注册用户路由
	user.Routers(r)
	// 注册产品路由
	product.Routers(r)
	// 注册库存路由
	inventory.Routers(r)
	return r
}
