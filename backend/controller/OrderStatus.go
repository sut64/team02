package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/entity"

)

// POST /order_statuses
func CreateOrderStatus(c *gin.Context) {
	var orderstatus entity.OrderStatus
	if err := c.ShouldBindJSON(&orderstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&orderstatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderstatus})
}

// GET /order_status/:id
func GetOrderStatus(c *gin.Context) {
	var orderstatus entity.OrderStatus
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM order_statuses WHERE id = ?", id).Scan(&orderstatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderstatus})
}

// GET /order_statuses
func ListOrderStatuses(c *gin.Context) {
	var orderstatuses []entity.OrderStatus
	if err := entity.DB().Raw("SELECT * FROM order_statuses").Scan(&orderstatuses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderstatuses})
}

// DELETE /order_statuses/:id
func DeleteOrderStatus(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM order_statuses WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "orderstatus not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /order_statuses
func UpdateOrderStatus(c *gin.Context) {
	var orderstatus entity.OrderStatus
	if err := c.ShouldBindJSON(&orderstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	if tx := entity.DB().Where("id = ?", orderstatus.ID).First(&orderstatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "orderstatus not found"})
		return
	}
 
	if err := entity.DB().Save(&orderstatus).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"data": orderstatus})
}
