package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Boontita/se-64-example/entity"

)

// POST /companies
func CreateCompany(c *gin.Context) {
	var company entity.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": company})
}

// GET /company/:id
func GetCompany(c *gin.Context) {
	var company entity.Company
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM companies WHERE id = ?", id).Scan(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": company})
}

// GET /companies
func ListCompanies(c *gin.Context) {
	var companies []entity.Company
	if err := entity.DB().Raw("SELECT * FROM companies").Scan(&companies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": companies})
}

// DELETE /companies/:id
func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM companies WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /companies
func UpdateCompany(c *gin.Context) {
	var company entity.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", company.ID).First(&company); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company not found"})
		return
	}

	if err := entity.DB().Save(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": company})
}
