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
	Sex      string `json:"sex" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserResponse struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	Sex        string `json:"sex" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type UpdateRoleRequest struct {
	Role  string `json:"role"`
	Sex   string `json:"sex"`
	Phone string `json:"phone"`
}

func GetUsers(c *gin.Context) {
	var users []user.User

	/*
		SELECT * FROM users
	*/
	if err := global.GL_DB.Model(&user.User{}).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}

	var resp []GetUserResponse
	for _, u := range users {
		resp = append(resp, GetUserResponse{
			Username:   u.Username,
			Role:       u.Role,
			Sex:        u.Sex,
			Phone:      u.Phone,
			CreateTime: u.CreateTime.UTC().Format("2006-01-02 15:04:05"),
			UpdateTime: u.UpdateTime.UTC().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func GetUserByName(c *gin.Context) {
	name := c.Param("name")

	var u user.User
	/*
		SELECT * FROM users WHERE username = name
	*/
	if err := global.GL_DB.Model(&user.User{}).Where("username = ?", name).First(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}

	resp := GetUserResponse{
		Username:   u.Username,
		Role:       u.Role,
		Sex:        u.Sex,
		Phone:      u.Phone,
		CreateTime: u.CreateTime.UTC().Format("2006-01-02 15:04:05"),
		UpdateTime: u.UpdateTime.UTC().Format("2006-01-02 15:04:05"),
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
		Sex:      request.Sex,
		Phone:    request.Phone,
	}

	/*
		INSERT INTO users (username, password, role, sex, phone, createtime, updatetime)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	*/
	if global.GL_DB.Model(&user.User{}).Create(&u).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func DeleteUserByName(c *gin.Context) {
	name := c.Param("name")

	var u user.User
	/*
		SELECT * FROM users WHERE username = name
		LIMIT 1
	*/
	if err := global.GL_DB.Model(&user.User{}).Where("username = ?", name).First(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
		return
	}

	if u.Role == "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete admin user"})
		return
	}

	/*
		DELETE FROM users WHERE username = name
	*/
	if global.GL_DB.Model(&user.User{}).Where("username = ?", name).Delete(&user.User{}).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var u user.User
	/*
		SELECT * FROM users WHERE username = request.Username
		LIMIT 1
	*/
	if err := global.GL_DB.Model(&user.User{}).Where("username = ?", request.Username).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}

	if !util.CheckPasswordHash(request.Password, u.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "role": u.Role})
}

func Logout(c *gin.Context) {
	// 自动备份数据库
	util.Backup(global.GL_VIPER)
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

func UpdateRoleByName(c *gin.Context) {
	name := c.Param("name")

	var u user.User
	/*
		SELECT * FROM users WHERE username = name
		LIMIT 1
	*/
	if err := global.GL_DB.Model(&user.User{}).Where("username = ?", name).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var request UpdateRoleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if u.Role == request.Role && u.Sex == request.Sex && u.Phone == request.Phone {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not changed"})
		return
	}

	// 不允许修改admin的role
	if request.Role != "" && u.Username == "admin" && request.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not allowed to change admin's role"})
		return
	}

	if request.Role != "" {
		u.Role = request.Role
	}
	if request.Phone != "" {
		u.Phone = request.Phone
	}
	if request.Sex != "" {
		u.Sex = request.Sex
	}

	// 使用Save方法，才会调用gorm的hooks
	/*
		Update users set role = ?, phone = ?, sex = ?
		where username = ?
	*/
	if global.GL_DB.Save(&u).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

func UpdatePasswordByName(c *gin.Context) {
	name := c.Param("name")

	var u user.User
	/*
		SELECT * FROM users WHERE username = name
		LIMIT 1
	*/
	if err := global.GL_DB.Model(&user.User{}).Where("username = ?", name).First(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	var request UpdatePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if util.CheckPasswordHash(request.Password, u.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not changed"})
		return
	}

	hashedPassword, err := util.HashAndSalt(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	if u.Password == hashedPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not changed"})
		return
	}

	u.Password = hashedPassword

	/*
		Update users set password = ?
		where username = ?
	*/
	if global.GL_DB.Save(&u).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}
