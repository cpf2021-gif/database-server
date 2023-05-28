package user

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.POST("/login", Login)

	e.POST("/users", CreateUser)
	e.GET("/users", GetUsers)
	e.PATCH("/users/:name", UpdateRoleByName)
}
