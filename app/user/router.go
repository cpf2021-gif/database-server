package user

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.POST("/login", Login)
	e.GET("/logout", Logout)

	e.POST("/users", CreateUser)
	e.GET("/users", GetUsers)
	e.GET("/users/:name", GetUserByName)
	e.PATCH("/users/:name", UpdateRoleByName)
	e.PATCH("/users/password/:name", UpdatePasswordByName)
	e.DELETE("/users/:name", DeleteUserByName)
}
