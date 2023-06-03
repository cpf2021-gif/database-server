package inventory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"server/global"
	"server/model/inventory"
)

type GetInventoriesResponse struct {
	ID          int    `json:"id" binding:"required"`
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
			ID:          ivt.ID,
			ProductName: ivt.ProductName,
			Quantity:    ivt.Quantity,
			MaxQuantity: ivt.MaxQuantity,
			MinQuantity: ivt.MinQuantity,
			CreateTime:  ivt.CreateTime.UTC().Format("2006-01-02 15:04:05"),
			UpdateTime:  ivt.UpdateTime.UTC().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

type GetInBoundsResponse struct {
	ID          int    `json:"id" binding:"required"`
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
			ID:          inb.ID,
			ProductName: inb.ProductName,
			Quantity:    inb.Quantity,
			UserName:    inb.UserName,
			CreateTime:  inb.CreateTime.UTC().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

type GetOutBoundsResponse struct {
	ID          int    `json:"id" binding:"required"`
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
			ID:          outb.ID,
			ProductName: outb.ProductName,
			Quantity:    outb.Quantity,
			UserName:    outb.UserName,
			CreateTime:  outb.CreateTime.UTC().Format("2006-01-02 15:04:05"),
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
	c.JSON(http.StatusOK, gin.H{"message": "success"})
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
		c.JSON(http.StatusOK, gin.H{"error": "failed to create outbound"})
		return
	}

	if ivt.Quantity < ivt.MinQuantity {
		c.JSON((http.StatusOK), gin.H{"message": "success, but inventory is less than min quantity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

type UpdateInventoryRequest struct {
	MAXQuantity int `json:"max_quantity" `
	MINQuantity int `json:"min_quantity" `
}

func UpdateInventoryRequestByID(c *gin.Context) {
	id := c.Param("id")
	var request UpdateInventoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ivt inventory.Inventory
	if err := global.GL_DB.Model(&inventory.Inventory{}).Where("id = ?", id).First(&ivt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get inventory"})
		return
	}

	if request.MAXQuantity != 0 {
		ivt.MaxQuantity = request.MAXQuantity
	}
	if request.MINQuantity != 0 {
		ivt.MinQuantity = request.MINQuantity
	}

	if ivt.MaxQuantity < ivt.MinQuantity {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "max quantity must be greater than min quantity"})
		return
	}

	if err := global.GL_DB.Save(&ivt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

type exportInboundRequest struct {
	ProductName string `json:"product_name"`
	Month       string `json:"month"`
}

func ExportInbound(c *gin.Context) {
	var request exportInboundRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var inbounds []inventory.Inbound
	if request.ProductName == "" && request.Month == "" {
		if err := global.GL_DB.Model(&inventory.Inbound{}).Find(&inbounds).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get inbounds"})
			return
		}
	} else {
		var db *gorm.DB
		db = global.GL_DB.Model(&inventory.Inbound{})
		if request.ProductName != "" {
			db = db.Where("product_name = ?", request.ProductName)
		}
		if request.Month != "" {
			db = db.Where("to_char(create_time, 'YYYY-MM') = ?", request.Month)
		}
		if err := db.Find(&inbounds).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get inbounds"})
			return
		}
	}

	var response []GetInBoundsResponse
	for _, inb := range inbounds {
		response = append(response, GetInBoundsResponse{
			ID:          inb.ID,
			ProductName: inb.ProductName,
			Quantity:    inb.Quantity,
			UserName:    inb.UserName,
			CreateTime:  inb.CreateTime.UTC().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

type exportOutboundRequest struct {
	ProductName string `json:"product_name"`
	Month       string `json:"month"`
}

func ExportOutbound(c *gin.Context) {
	var request exportOutboundRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var outbounds []inventory.Outbound
	if request.ProductName == "" && request.Month == "" {
		if err := global.GL_DB.Model(&inventory.Outbound{}).Find(&outbounds).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get outbounds"})
			return
		}
	} else {
		var db *gorm.DB
		db = global.GL_DB.Model(&inventory.Outbound{})
		if request.ProductName != "" {
			db = db.Where("product_name = ?", request.ProductName)
		}
		if request.Month != "" {
			db = db.Where("to_char(create_time, 'YYYY-MM') = ?", request.Month)
		}
		if err := db.Find(&outbounds).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get outbounds"})
			return
		}
	}

	var response []GetOutBoundsResponse
	for _, outb := range outbounds {
		response = append(response, GetOutBoundsResponse{
			ID:          outb.ID,
			ProductName: outb.ProductName,
			Quantity:    outb.Quantity,
			UserName:    outb.UserName,
			CreateTime:  outb.CreateTime.UTC().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
