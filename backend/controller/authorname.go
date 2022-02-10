package controller
import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/entity"
)

// POST /authornames
func CreateAuthorName(c *gin.Context) {
	var authorname entity.AuthorName
	if err := c.ShouldBindJSON(&authorname); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&authorname).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": authorname})
}

// GET /authorname/:id
func GetAuthorName(c *gin.Context) {
	var authorname entity.AuthorName
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM author_names WHERE id = ?", id).Scan(&authorname).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": authorname})
}

// GET /authornames
func ListAuthorName(c *gin.Context) {
	var authornames []entity.AuthorName
	if err := entity.DB().Raw("SELECT * FROM author_names").Scan(&authornames).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": authornames})
}

// DELETE /authornames/:id
func DeleteAuthorName(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM author_names WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorname not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /authornames
func UpdateAuthorName(c *gin.Context) {
	var authorname entity.AuthorName
	if err := c.ShouldBindJSON(&authorname); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", authorname.ID).First(&authorname); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorname not found"})
		return
	}

	if err := entity.DB().Save(&authorname).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": authorname})
}