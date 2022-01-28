package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PloychompooWongchanthat/se-64/entity"

)

// POST /book_types
func CreateBookType(c *gin.Context) {
	var booktype entity.BookType
	if err := c.ShouldBindJSON(&booktype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&booktype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": booktype})
}

// GET /booktype/:id
func GetBookType(c *gin.Context) {
	var booktype entity.BookType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM book_types WHERE id = ?", id).Scan(&booktype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booktype})
}

// GET /book_types
func ListBookTypes(c *gin.Context) {
	var booktypes []entity.BookType
	if err := entity.DB().Raw("SELECT * FROM book_types").Scan(&booktypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booktypes})
}

// DELETE /book_types/:id
func DeleteBookType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM book_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booktype not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /book_types
func UpdateBookType(c *gin.Context) {
	var booktype entity.BookType
	if err := c.ShouldBindJSON(&booktype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", booktype.ID).First(&booktype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booklocation not found"})
		return
	}

	if err := entity.DB().Save(&booktype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booktype})
}