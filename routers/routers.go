package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"server/app/inventory"
	"server/app/product"
	"server/app/user"
	"server/global"
	"server/util"

	dataInventory "server/model/inventory"
	dataProduct "server/model/product"
	dataUser "server/model/user"
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

	// 获取整体数据
	r.GET("/data", func(c *gin.Context) {
		var usersnumber int64
		var productsnumber int64
		var inventorynumber int64
		var inboundnumber int64
		var outboundnumber int64
		var days int64

		// 获取用户数量
		global.GL_DB.Model(&dataUser.User{}).Count(&usersnumber)
		// 获取产品数量
		global.GL_DB.Model(&dataProduct.Product{}).Count(&productsnumber)
		// 获取库存数量
		global.GL_DB.Model(&dataInventory.Inventory{}).Select("sum(quantity)").Row().Scan(&inventorynumber)
		// 获取本月入库数量
		global.GL_DB.Model(&dataInventory.Inbound{}).Where("DATE_TRUNC('month',create_time) = DATE_TRUNC('month', current_date at time zone 'Asia/Shanghai')").Select("sum(quantity)").Row().Scan(&inboundnumber)
		// 获取本月出库数量
		global.GL_DB.Model(&dataInventory.Outbound{}).Where("DATE_TRUNC('month',create_time) = DATE_TRUNC('month', current_date at time zone 'Asia/Shanghai')").Select("sum(quantity)").Row().Scan(&outboundnumber)
		// 获取天数
		// "2023-10-11" - "2023-10-10"
		startDate, _ := time.Parse("2006-01-02", global.GL_CONFIG.App.CreatedTime)
		endDate, _ := time.Parse("2006-01-02", global.GL_CONFIG.Database.BackupTime)

		duration := endDate.Sub(startDate)
		days = (int64)(duration.Hours()/24) + 1

		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"usersnumber":     usersnumber,
				"productsnumber":  productsnumber,
				"inventorynumber": inventorynumber,
				"inboundnumber":   inboundnumber,
				"outboundnumber":  outboundnumber,
				"days":            days,
			},
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
