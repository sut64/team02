package controller

import (
	//"time"

	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/entity"
)

// POST /book_informations
func CreateBookInformation(c *gin.Context) {
	//now := time.Now()
	var bookinformation entity.BookInformation
	var booklocation entity.BookLocation
	var bookorder entity.BookOrder
	var bookcategory entity.BookCategory
	var member 		entity.Member

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร bookinformation
	if err := c.ShouldBindJSON(&bookinformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา member ด้วย id
	if tx := entity.DB().Where("id = ?", bookinformation.MemberID).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	// ค้นหา booklocation ด้วย id
	if tx := entity.DB().Where("id = ?", bookinformation.BookLocationID).First(&booklocation); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booklocation not found"})
		return
	}

	// ค้นหา bookorder ด้วย id
	if tx := entity.DB().Where("id = ?", bookinformation.BookOrderID).First(&bookorder); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookorder not found"})
		return
	}

	// ค้นหา bookcategory ด้วย id
	if tx := entity.DB().Where("id = ?", bookinformation.BookCategoryID).First(&bookcategory); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookcategory not found"})
		return
	}

	//สร้าง bookinformation
	bi := entity.BookInformation{
		Member: 		 member,						//โยงความสัมพันธ์กับ Entity member
		BookLocation:    booklocation,                  // โยงความสัมพันธ์กับ Entity BookLocation
		BookOrder:       bookorder,                     // โยงความสัมพันธ์กับ Entity BookOrder
		BookCategory:    bookcategory,                  // โยงความสัมพันธ์กับ Entity BookCategory
		Date:            bookinformation.Date,            //ตั้งค่าฟิลก์ Date
		YearPublication: bookinformation.YearPublication, //ตั้งค่าฟิลก์ YearPublication
		CallNumber:      bookinformation.CallNumber,      //ตั้งค่าฟิลก์ CallNumber
	}

	//แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(bi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// บันทึก
	if err := entity.DB().Create(&bi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bi})
}

// GET /bookinformation/:id
func GetBookInformation(c *gin.Context) {
	var bookinformation entity.BookInformation
	id := c.Param("id")
	if err := entity.DB().Preload("BookLocation").Preload("BookOrder").Preload("BookCategory").Raw("SELECT * FROM book_informations WHERE id = ?", id).Find(&bookinformation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookinformation})
}

// GET /book_informations
func ListBookInformations(c *gin.Context) {
	var bookinformations []entity.BookInformation
	if err := entity.DB().Preload("BookLocation").Preload("BookOrder").Preload("BookCategory").Raw("SELECT * FROM book_informations").Find(&bookinformations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookinformations})
}

// DELETE /book_informations/:id
func DeleteBookInformation(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM book_informations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookinformation not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /book_informations
func UpdateBookInformation(c *gin.Context) {
	var bookinformation entity.BookInformation
	if err := c.ShouldBindJSON(&bookinformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bookinformation.ID).First(&bookinformation); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookinformation not found"})
		return
	}

	if err := entity.DB().Save(&bookinformation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookinformation})
}
