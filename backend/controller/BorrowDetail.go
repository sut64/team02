package controller

import (
	"github.com/sut64/team02/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /borrows
func CreateBorrow(c *gin.Context) {

	var borrowDetail entity.BorrowDetail
	var serviceplace entity.ServicePlace
	var info entity.BookInformation
	var member entity.Member
	var status entity.Status

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร borrowDetail
	if err := c.ShouldBindJSON(&borrowDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา servicePlace ด้วย id
	if tx := entity.DB().Where("id = ?", borrowDetail.ServicePlaceID).First(&serviceplace); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "servicePlace not found"})
		return
	}

	// 10: ค้นหา BookInfo ด้วย id
	if tx := entity.DB().Where("id = ?", borrowDetail.InfoID).First(&info); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Infomation not found"})
		return
	}

	// 12: ค้นหา member ด้วย id
	if tx := entity.DB().Where("id = ?", borrowDetail.MemberID).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	// 13: ค้นหา status ด้วย id
	if tx := entity.DB().Where("id = ?", borrowDetail.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	// 14: สร้าง BorrowDetail
	bd := entity.BorrowDetail{
		Member:         member,
		ServicePlace:   serviceplace,
		Info:           info,
		Status:         status,
		DateToBorrow:   borrowDetail.DateToBorrow,
		Tel:            borrowDetail.Tel,
		BorrowDuration: borrowDetail.BorrowDuration,
	}

	/*//validate
	if _, err := govalidator.ValidateStruct(bd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}*/

	// 15: บันทึก
	if err := entity.DB().Create(&bd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bd})
}

// GET /borrow/:id
func GetBorrow(c *gin.Context) {
	var borrowDetail entity.BorrowDetail
	id := c.Param("id")
	if err := entity.DB().Preload("ServicePlace").Preload("Status").Preload("Info").Preload("Member").Raw("SELECT * FROM borrow_details WHERE id = ?", id).Find(&borrowDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": borrowDetail})
}

// GET /borrows
func ListBorrows(c *gin.Context) {
	var borrowDetail []entity.BorrowDetail
	if err := entity.DB().Preload("ServicePlace").Preload("Status").Preload("Member").Preload("Info").Preload("Info.BookOrder").Raw("SELECT * FROM borrow_details").Find(&borrowDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": borrowDetail})
}

// GET /borrow/member/:id
func ListBorrowsByMember(c *gin.Context) {
	var borrowDetail []entity.BorrowDetail
	id := c.Param("id")
	if err := entity.DB().Preload("ServicePlace").Preload("Status").Preload("Info").Preload("Member").Raw("SELECT * FROM borrow_details WHERE member_id = ?", id).Find(&borrowDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": borrowDetail})
}

// DELETE /borrows/:id
func DeleteBorrow(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM borrow_details WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "active pro not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /borrows
func UpdateBorrow(c *gin.Context) {
	var borrowDetail entity.BorrowDetail
	if err := c.ShouldBindJSON(&borrowDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", borrowDetail.ID).First(&borrowDetail); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "active not found"})
		return
	}

	if err := entity.DB().Save(&borrowDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": borrowDetail})
}
