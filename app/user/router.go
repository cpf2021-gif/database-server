package user

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.POST("/login", Login)

	e.POST("/users", CreateUser)
	e.GET("/users/:id", GetUserByID)
	e.GET("/users", GetUsers)
	e.DELETE("/users/:id", DeleteUserByID)
	e.PATCH("/users/:id", UpdateRoleByID)
}
