package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team02/entity"
)

// POST /members
func CreateMember(c *gin.Context) {
	var member entity.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": member})
}

// GET /member/:id
func GetMember(c *gin.Context) {
	var member entity.Member
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM members WHERE id = ?", id).Scan(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}

// GET /members
func ListMembers(c *gin.Context) {
	var members []entity.Member
	if err := entity.DB().Raw("SELECT * FROM members").Scan(&members).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": members})
}

// DELETE /members/:id
func DeleteMember(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM members WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /members
func UpdateMember(c *gin.Context) {
	var member entity.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", member.ID).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repairstatus not found"})
		return
	}

	if err := entity.DB().Save(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}