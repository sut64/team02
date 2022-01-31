package controller

import (
	"github.com/sut64/team02/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /bookReturns
func CreateBookReturn(c *gin.Context) {

	var book_return entity.BookReturn
	var member entity.Member
	var borrowDetail entity.BorrowDetail
	var serviceplace entity.ServicePlace
	var status entity.Status

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร bookReturn
	if err := c.ShouldBindJSON(&book_return); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 9: ค้นหา borrowDetail ด้วย id
	if tx := entity.DB().Where("id = ?", book_return.BorrowDetailID).First(&borrowDetail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "borrowdetail not found"})
		return
	}

	// 10: ค้นหา serviceplace ด้วย id
	if tx := entity.DB().Where("id = ?", book_return.ServicePlaceID).First(&serviceplace); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "seviceplace not found"})
		return
	}

	// 11: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", book_return.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	// 12: สร้าง BookReturn
	wv := entity.BookReturn{
		Member:       member,
		BorrowDetail: borrowDetail,
		ServicePlace: serviceplace,
		Status:       status,
		Damage:       book_return.Damage,
		Tel:          book_return.Tel,
		DateReturn:   book_return.DateReturn,
	}

	// // แทรกการ validate ไว้ช่วงนี้ของ controller
	// if _, err := govalidator.ValidateStruct(book_return); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /book_return/:id

func GetBookReturn(c *gin.Context) {

	var book_return entity.BookReturn

	id := c.Param("id")

	if err := entity.DB().Preload("BorrowDetail").Preload("SevicePlace").Preload("Status").Raw("SELECT * FROM book_returns WHERE id = ?", id).Scan(&book_return).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": book_return})

}

// GET /book_returns

func ListBookReturns(c *gin.Context) {

	var book_returns []entity.BookReturn

	if err := entity.DB().Preload("BorrowDetail").Preload("SevicePlace").Preload("Status").Raw("SELECT * FROM book_returns").Scan(&book_returns).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": book_returns})

}

// DELETE /book_returns/:id

func DeleteBookReturn(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM book_returns WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "book_return not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /book_returns

func UpdateBookReturn(c *gin.Context) {

	var book_return entity.BookReturn

	if err := c.ShouldBindJSON(&book_return); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", book_return.ID).First(&book_return); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "book_return not found"})

		return

	}

	if err := entity.DB().Save(&book_return).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": book_return})

}
