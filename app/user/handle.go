package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/global"
	"server/model/user"
	"server/util"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UpdateRoleRequest struct {
	Role string `json:"role" binding:"required"`
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var u user.User
	if err := global.GL_DB.Model(&user.User{}).Where("id = ?", id).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	resp := GetUserResponse{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role,
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func GetUsers(c *gin.Context) {
	var users []user.User

	if err := global.GL_DB.Model(&user.User{}).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}

	var resp []GetUserResponse
	for _, u := range users {
		resp = append(resp, GetUserResponse{
			ID:       u.ID,
			Username: u.Username,
			Role:     u.Role,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func CreateUser(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashAndSalt(request.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	u := user.User{
		Username: request.Username,
		Password: hashedPassword,
		Role:     request.Role,
	}

	if global.GL_DB.Model(&user.User{}).Create(&u).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var u user.User
	if err := global.GL_DB.Model(&user.User{}).Where("username = ?", request.Username).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}

	if !util.CheckPasswordHash(request.Password, u.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")

	var u user.User
	if err := global.GL_DB.Model(&user.User{}).Where("id = ?", id).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if global.GL_DB.Model(&user.User{}).Delete(&u).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func UpdateRoleByID(c *gin.Context) {
	id := c.Param("id")

	var u user.User
	if err := global.GL_DB.Model(&user.User{}).Where("id = ?", id).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var request UpdateRoleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Role = request.Role

	if global.GL_DB.Model(&user.User{}).Save(&u).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}
