package controller
import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/entity"
)

// POST /typeresearches
func CreateTypeResearch(c *gin.Context) {
	var typeresearch entity.TypeResearch
	if err := c.ShouldBindJSON(&typeresearch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&typeresearch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": typeresearch})
}

// GET /typeresearch/:id
func GetTypeResearch(c *gin.Context) {
	var typeresearch entity.TypeResearch
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM type_researches WHERE id = ?", id).Scan(&typeresearch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": typeresearch})
}

// GET /typeresearches
func ListTypeResearch(c *gin.Context) {
	var typeresearches []entity.TypeResearch
	if err := entity.DB().Raw("SELECT * FROM type_researches").Scan(&typeresearches).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": typeresearches})
}

// DELETE /typeresearches/:id
func DeleteTypeResearch(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM type_researches WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeresearch not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /typeresearches
func UpdateTypeResearch(c *gin.Context) {
	var typeresearch entity.TypeResearch
	if err := c.ShouldBindJSON(&typeresearch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", typeresearch.ID).First(&typeresearch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeresearch not found"})
		return
	}

	if err := entity.DB().Save(&typeresearch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": typeresearch})
}
