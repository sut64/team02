package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/entity"
)

// POST /researches
func CreateResearch(c *gin.Context) {

	var member entity.Member
	var research entity.Research
	var typeresearch entity.TypeResearch
	var authorname entity.AuthorName
	var institutionname entity.InstitutionName

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร researche
	if err := c.ShouldBindJSON(&research); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา typeresearch ด้วย id
	if tx := entity.DB().Where("id = ?", research.TypeResearchID).First(&typeresearch); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeresearch not found"})
		return
	}

	// 10: ค้นหา authorname ด้วย id
	if tx := entity.DB().Where("id = ?", research.AuthorNameID).First(&authorname); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorname not found"})
		return
	}

	// 11: ค้นหา institutionname ด้วย id
	if tx := entity.DB().Where("id = ?", research.InstitutionNameID).First(&institutionname); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "institutionname not found"})
		return
	}
	// 12: สร้าง Research
	rs := entity.Research{
		Member:            member,
		TypeResearch:      typeresearch,               // โยงความสัมพันธ์กับ Entity TypeResearch
		AuthorName:        authorname,                 // โยงความสัมพันธ์กับ Entity AuthorName
		InstitutionName:   institutionname,            // โยงความสัมพันธ์กับ Entity InstitutionName
		NameResearch:      research.NameResearch,      //ตั้งค่าฟิลด์ NameResearch
		YearOfPublication: research.YearOfPublication, //ตั้งค่าฟิลด์ YearOfPublication
		RecordingDate:     research.RecordingDate,     //ตั้งค่าฟิลด์ RecordingDate
	}
	//validate
	if _, err := govalidator.ValidateStruct(rs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&rs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rs})
}

// GET /research/:id
func GetResearch(c *gin.Context) {
	var research entity.Research
	id := c.Param("id")
	if err := entity.DB().Preload("TypeResearch").Preload("AuthorName").Preload("InstitutionName").Raw("SELECT * FROM researches WHERE id = ?", id).Find(&research).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": research})
}

// GET /researches
func ListResearch(c *gin.Context) {
	var researches []entity.Research
	if err := entity.DB().Preload("TypeResearch").Preload("AuthorName").Preload("InstitutionName").Raw("SELECT * FROM researches").Find(&researches).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": researches})
}

// DELETE /researches/:id
func DeleteResearch(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM researches WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "research not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /researches
func UpdateResearch(c *gin.Context) {
	var research entity.Research
	if err := c.ShouldBindJSON(&research); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", research.ID).First(&research); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "research not found"})
		return
	}

	if err := entity.DB().Save(&research).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": research})
}
