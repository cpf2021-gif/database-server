package inventory

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/global"
	"server/model/inventory"
)

type GetInventoriesResponse struct {
	ProductName string `json:"product_name" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	MaxQuantity int    `json:"max_quantity" binding:"required"`
	MinQuantity int    `json:"min_quantity" binding:"required"`
	CreateTime  string `json:"create_time" binding:"required"`
	UpdateTime  string `json:"update_time" binding:"required"`
}

/*
SELECT * FROM inventories;
*/
func GetInventories(c *gin.Context) {
	var inventories []inventory.Inventory

	if err := global.GL_DB.Model(&inventory.Inventory{}).Find(&inventories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get inventories"})
		return
	}

	var response []GetInventoriesResponse
	for _, ivt := range inventories {
		response = append(response, GetInventoriesResponse{
			ProductName: ivt.ProductName,
			Quantity:    ivt.Quantity,
			MaxQuantity: ivt.MaxQuantity,
			MinQuantity: ivt.MinQuantity,
			CreateTime: ivt.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: ivt.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

type GetInBoundsResponse struct {
	ProductName string `json:"product_name" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	UserName    string `json:"user_name" binding:"required"`
	CreateTime  string `json:"create_time" binding:"required"`
}

/*
SELECT * FROM inbounds;
*/
func GetInBounds(c *gin.Context) {
	var inBounds []inventory.Inbound

	if err := global.GL_DB.Model(&inventory.Inbound{}).Find(&inBounds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get inbounds"})
		return
	}

	var response []GetInBoundsResponse
	for _, inb := range inBounds {
		response = append(response, GetInBoundsResponse{
			ProductName: inb.ProductName,
			Quantity:    inb.Quantity,
			UserName:    inb.UserName,
			CreateTime: inb.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

type GetOutBoundsResponse struct {
	ProductName string `json:"product_name" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	UserName    string `json:"user_name" binding:"required"`
	CreateTime  string `json:"create_time" binding:"required"`
}

/*
SELECT * FROM outbounds;
*/
func GetOutBounds(c *gin.Context) {
	var outBounds []inventory.Outbound

	if err := global.GL_DB.Model(&inventory.Outbound{}).Find(&outBounds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get outbounds"})
		return
	}

	var response []GetOutBoundsResponse
	for _, outb := range outBounds {
		response = append(response, GetOutBoundsResponse{
			ProductName: outb.ProductName,
			Quantity:    outb.Quantity,
			UserName:    outb.UserName,
			CreateTime: outb.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

type CreateInventoryRequest struct {
	ProductName string `json:"product_name" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	UserName    string `json:"user_name" binding:"required"`
}

func CreateInbound(c *gin.Context) {
	var request CreateInventoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ivt inventory.Inventory
	// 不存在则创建
	if err := global.GL_DB.Model(&inventory.Inventory{}).Where("product_name = ?", request.ProductName).First(&ivt).Error; err != nil {
		ivt.MaxQuantity = request.Quantity * 2
		ivt.ProductName = request.ProductName
		ivt.MinQuantity = request.Quantity / 2
		global.GL_DB.Create(&ivt)	
	}

	if err := global.GL_DB.Model(&inventory.Inventory{}).Where("product_name = ?", request.ProductName).First(&ivt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get inventory"})
		return
	}

	ivt.Quantity += request.Quantity
	if ivt.Quantity > ivt.MaxQuantity {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "exceed max quantity"})
		return
	}

	if err := global.GL_DB.Save(&ivt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update inventory"})
		return
	}

	inbound := inventory.Inbound{
		ProductName: request.ProductName,
		Quantity:    request.Quantity,
		UserName:    request.UserName,
	}
	if err := global.GL_DB.Create(&inbound).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create inbound"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": inbound})
}

type CreateOutboundRequest struct {
	ProductName string `json:"product_name" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	UserName    string `json:"user_name" binding:"required"`
}

func CreateOutbound(c *gin.Context) {
	var request CreateOutboundRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ivt inventory.Inventory
	if err := global.GL_DB.Model(&inventory.Inventory{}).Where("product_name = ?", request.ProductName).First(&ivt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get inventory"})
		return
	}

	if ivt.Quantity < request.Quantity {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not enough inventory"})
		return
	}

	ivt.Quantity -= request.Quantity
	if err := global.GL_DB.Save(&ivt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update inventory"})
		return
	}

	outbound := inventory.Outbound{
		ProductName: request.ProductName,
		Quantity:    request.Quantity,
		UserName:    request.UserName,
	}

	if err := global.GL_DB.Create(&outbound).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create outbound"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": outbound})
}
