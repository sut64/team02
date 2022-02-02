package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/entity"

)

// POST /book_categories
func CreateBookCategory(c *gin.Context) {
	var bookcategory entity.BookCategory
	if err := c.ShouldBindJSON(&bookcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&bookcategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookcategory})
}

// GET /bookcategory/:id
func GetBookCategory(c *gin.Context) {
	var bookcategory entity.BookCategory
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM book_categories WHERE id = ?", id).Scan(&bookcategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookcategory})
}

// GET /book_categories
func ListBookCategories(c *gin.Context) {
	var bookcategories []entity.BookCategory
	if err := entity.DB().Raw("SELECT * FROM book_categories").Scan(&bookcategories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookcategories})
}

// DELETE /book_categories/:id
func DeleteBookCategory(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM book_categories WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookcategory not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /book_categories
func UpdateBookCategory(c *gin.Context) {
	var bookcategory entity.BookCategory
	if err := c.ShouldBindJSON(&bookcategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bookcategory.ID).First(&bookcategory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookcategory not found"})
		return
	}

	if err := entity.DB().Save(&bookcategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookcategory})
}