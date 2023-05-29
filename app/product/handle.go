package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/global"
	"server/model/product"
)

type CreateProductRequest struct {
	Name         string `json:"name" binding:"required"`
	SupplierName string `json:"supplier_name" binding:"required"`
}

type UpdateProductRequest struct {
	SupplierName string `json:"supplier_name" binding:"required"`
}

// 创建商品
func CreateProduct(c *gin.Context) {
	var request CreateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := product.Product{
		Name:         request.Name,
		SupplierName: request.SupplierName,
	}

	/*
		INSERT INTO products (name, supplier_name, create_time, update_time)
		VALUES ('name', 'supplier_name', 'create_time', 'update_time')
	*/
	if err := global.GL_DB.Model(&product.Product{}).Create(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}

type GetProductsResponse struct {
	ID           int   `json:"id"`
	Name         string `json:"name"`
	SupplierName string `json:"supplier_name"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

// 获取所有商品
func GetProducts(c *gin.Context) {
	var products []product.Product

	/*
		SELECT * FROM products
	*/
	if err := global.GL_DB.Model(&product.Product{}).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products"})
		return
	}

	var response []GetProductsResponse
	for _, p := range products {
		response = append(response, GetProductsResponse{
			ID:           p.ID,
			Name:         p.Name,
			SupplierName: p.SupplierName,
			CreateTime:  p.CreateTime.UTC().Format("2006-01-02 15:04:05"),
			UpdateTime:  p.UpdateTime.UTC().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// 修改商品供应商
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var request UpdateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var p product.Product
	/*
		SELECT * FROM products WHERE id = ?
	*/
	if err := global.GL_DB.Model(&product.Product{}).Where("id = ?", id).First(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	if p.SupplierName == request.SupplierName {
		c.JSON(http.StatusBadRequest, gin.H{"error": "supplier name is the same"})
	}
	/*
		UPDATE products SET supplier_name = ?, update_time = ? WHERE id = ?
	*/
	if err := global.GL_DB.Model(&p).Updates(request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}

type CreateSupplierRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Location string `json:"location" binding:"required"`
}

// 创建供应商
func CreateSupplier(c *gin.Context) {
	var request CreateSupplierRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := product.Supplier{
		Name:     request.Name,
		Phone:    request.Phone,
		Location: request.Location,
	}

	/*
		INSERT INTO suppliers (name, phone, location, create_time, update_time)
		VALUES ('name', 'phone', 'location', 'create_time', 'update_time')
	*/
	if err := global.GL_DB.Model(&product.Supplier{}).Create(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": s})
}
