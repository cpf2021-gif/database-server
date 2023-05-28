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

// 定义跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 中间件
	r.Use(Cors())
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
