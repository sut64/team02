package controller
import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/entity"
)

// POST /institutionnames
func CreateInstitutionName(c *gin.Context) {
	var institutionname entity.InstitutionName
	if err := c.ShouldBindJSON(&institutionname); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&institutionname).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": institutionname})
}

// GET /institutionname/:id
func GetInstitutionName(c *gin.Context) {
	var institutionname entity.InstitutionName
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM institution_names WHERE id = ?", id).Scan(&institutionname).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institutionname})
}

// GET /institutionnames
func ListTypeInstitutionName(c *gin.Context) {
	var institutionnames []entity.InstitutionName
	if err := entity.DB().Raw("SELECT * FROM institution_names").Scan(&institutionnames).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institutionnames})
}

// DELETE /institutionnames/:id
func DeleteInstitutionName(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM institution_names WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "institutionname not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /institutionnames
func UpdateInstitutionName(c *gin.Context) {
	var institutionname entity.InstitutionName
	if err := c.ShouldBindJSON(&institutionname); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", institutionname.ID).First(&institutionname); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "institutionname not found"})
		return
	}

	if err := entity.DB().Save(&institutionname).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institutionname})
}