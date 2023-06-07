package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/global"
	"server/model/inventory"
	"server/model/product"
)

type CreateProductRequest struct {
	Name string `json:"name" binding:"required"`
}

// 创建商品
func CreateProduct(c *gin.Context) {
	var request CreateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := product.Product{
		Name: request.Name,
	}

	/*
		INSERT INTO products (name , create_time, update_time)
		VALUES ('name', 'create_time', 'update_time')
	*/
	if err := global.GL_DB.Model(&product.Product{}).Create(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 查看库存中是否已经存在该商品，不存在则创建
	var ivt inventory.Inventory
	if err := global.GL_DB.Model(&inventory.Inventory{}).Where("product_name = ?", p.Name).First(&ivt).Error; err != nil {
		if err := global.GL_DB.Model(&inventory.Inventory{}).Create(&inventory.Inventory{
			ProductName: p.Name,
			Quantity:    0,
			MinQuantity: 10,
			MaxQuantity: 100,
		}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "create product success"})
}

type GetProductsResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
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
			ID:         p.ID,
			Name:       p.Name,
			CreateTime: p.CreateTime.UTC().Format("2006-01-02 15:04:05"),
			UpdateTime: p.UpdateTime.UTC().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// 获取所有的供应商
func GetSuppliers(c *gin.Context) {
	var suppliers []product.Supplier

	/*
		SELECT * FROM suppliers
	*/
	if err := global.GL_DB.Model(&product.Supplier{}).Find(&suppliers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get suppliers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": suppliers})
}

// 获取所有销售商
func GetSellers(c *gin.Context) {
	var sellers []product.Seller

	/*
		SELECT * FROM sellers
	*/
	if err := global.GL_DB.Model(&product.Seller{}).Find(&sellers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get sellers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sellers})
}

type UpdateProductRequest struct {
	Name string `json:"name" binding:"required"`
}

// 修改商品名字
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

	if request.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name can not be empty"})
		return
	}

	if request.Name == p.Name {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name can not be same"})
		return
	}

	/*
		UPDATE products SET name = ? , update_time = ? WHERE id = ?
	*/
	if err := global.GL_DB.Save(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update product success"})
}

type CreateSupplierRequest struct {
	Name string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
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
		Name: request.Name,
		Phone: request.Phone,
		Location: request.Location,
	}

	/*
		INSERT INTO suppliers (name , phone, location)
		VALUES ('name', 'phone', 'location')
	*/
	if err := global.GL_DB.Model(&product.Supplier{}).Create(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create supplier success"})
}

type CreateSellerRequest struct {
	Name string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Location string `json:"location" binding:"required"`
}

// 创建销售商
func CreateSeller(c *gin.Context) {
	var request CreateSellerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := product.Seller{
		Name: request.Name,
		Phone: request.Phone,
		Location: request.Location,
	}

	/*
		INSERT INTO sellers (name , phone, location)
		VALUES ('name', 'phone', 'location')
	*/
	if err := global.GL_DB.Model(&product.Seller{}).Create(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create seller success"})
}

// 修改商品的供应商
type UpdateSupplierRequest struct {
	Location string `json:"location"`
	Phone string `json:"phone"`
}

func UpdateSupplier(c *gin.Context) {
	name := c.Param("name")

	var request UpdateSupplierRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var s product.Supplier
	/*
		SELECT * FROM suppliers WHERE name = ?
	*/
	if err := global.GL_DB.Model(&product.Supplier{}).Where("name = ?", name).First(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update supplier"})
		return
	}

	if request.Location == "" && request.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location and phone can not be empty"})
		return
	}

	if request.Location == s.Location && request.Phone == s.Phone {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location and phone can not be same"})
		return
	}

	if request.Location != ""  && request.Location != s.Location {
		s.Location = request.Location
	}

	if request.Phone != "" && request.Phone != s.Phone {
		s.Phone = request.Phone
	}

	/*
		UPDATE suppliers SET location = ? , phone = ? 
		WHERE name = ?
	*/

	if err := global.GL_DB.Save(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update supplier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update supplier success"})
}

// 修改商品的销售商
type UpdateSellerRequest struct {
	Location string `json:"location"`
	Phone string `json:"phone"`
}

func UpdateSeller(c *gin.Context) {
	name := c.Param("name")

	var request UpdateSellerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var s product.Seller
	/*
		SELECT * FROM sellers WHERE name = ?
	*/
	if err := global.GL_DB.Model(&product.Seller{}).Where("name = ?", name).First(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update seller"})
		return
	}

	if request.Location == "" && request.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location and phone can not be empty"})
		return
	}

	if request.Location == s.Location && request.Phone == s.Phone {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location and phone can not be same"})
		return
	}

	if request.Location != ""  && request.Location != s.Location {
		s.Location = request.Location
	}

	if request.Phone != "" && request.Phone != s.Phone {
		s.Phone = request.Phone
	}

	/*
		UPDATE sellers SET location = ? , phone = ? 
		WHERE
	*/
	if err := global.GL_DB.Save(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update seller"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update seller success"})
}