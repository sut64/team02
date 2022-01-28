package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/entity"

)

// POST /book_locations
func CreateBookLocation(c *gin.Context) {
	var booklocation entity.BookLocation
	if err := c.ShouldBindJSON(&booklocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&booklocation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": booklocation})
}

// GET /booklocation/:id
func GetBookLocation(c *gin.Context) {
	var booklocation entity.BookLocation
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM book_locations WHERE id = ?", id).Scan(&booklocation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booklocation})
}

// GET /book_locations
func ListBookLocations(c *gin.Context) {
	var booklocations []entity.BookLocation
	if err := entity.DB().Raw("SELECT * FROM book_locations").Scan(&booklocations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booklocations})
}

// DELETE /book_locations/:id
func DeleteBookLocation(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM book_locations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booklocation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /book_locations
func UpdateBookLocation(c *gin.Context) {
	var booklocation entity.BookLocation
	if err := c.ShouldBindJSON(&booklocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", booklocation.ID).First(&booklocation); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booklocation not found"})
		return
	}

	if err := entity.DB().Save(&booklocation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booklocation})
}