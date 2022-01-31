package controller

import (
	"github.com/sut64/team02/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /places

func ListPlaces(c *gin.Context) {

	var fests []entity.ServicePlace

	if err := entity.DB().Raw("SELECT * FROM service_places").Scan(&fests).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": fests})

}

// GET /place/:id

func GetPlace(c *gin.Context) {

	var fest entity.ServicePlace

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM service_places WHERE id = ?", id).Scan(&fest).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": fest})

}

// POST /places

func CreatePlace(c *gin.Context) {

	var fest entity.ServicePlace

	if err := c.ShouldBindJSON(&fest); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&fest).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": fest})

}

// PATCH /places

func UpdatePlace(c *gin.Context) {

	var fest entity.ServicePlace

	if err := c.ShouldBindJSON(&fest); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", fest.ID).First(&fest); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "servicePlace not found"})

		return

	}

	if err := entity.DB().Save(&fest).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": fest})

}

// DELETE /places/:id

func DeletePlace(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM service_places WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "servicePlace not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}
