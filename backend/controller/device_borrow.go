package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team02/entity"
	"github.com/asaskevich/govalidator"
)


// POST /deviceborrows
func CreateDeviceBorrow(c *gin.Context) {
	
	var deviceborrow entity.DeviceBorrow
	var member 		entity.Member
	var devicelist entity.DeviceList
	var devicetype entity.DeviceType

	// ผลลัพธ์ที่ได้จากขั้นตอนที่  จะถูก bind เข้าตัวแปร deviceborrow
	if err := c.ShouldBindJSON(&deviceborrow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา member ด้วย id
	if tx := entity.DB().Where("id = ?", deviceborrow.MemberID).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	// ค้นหา devicelist ด้วย id
	if tx := entity.DB().Where("id = ?", deviceborrow.DeviceListID).First(&devicelist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "devicelist not found"})
		return
	}

	// ค้นหา devicetype ด้วย id
	if tx := entity.DB().Where("id = ?", deviceborrow.DeviceTypeID).First(&devicetype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "devicetype not found"})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(deviceborrow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//สร้าง deviceborrow
	bd := entity.DeviceBorrow{
		Member: 			member,						//โยงความสัมพันธ์กับ Entity member
		DeviceList:			devicelist,					//โยงความสัมพันธ์กับ Entity devicelist
		DeviceType:			devicetype,					//โยงความสัมพันธ์กับ Entity devicetype
		DeviceName:			deviceborrow.DeviceName,	//ตั้งค่าฟิลก์ DeviceName
		BorrowCode:			deviceborrow.BorrowCode,	//ตั้งค่าฟิลก์ BorrowCode
		Amount:				deviceborrow.Amount,		//ตั้งค่าฟิลก์ Amount
		Date:				deviceborrow.Date,			//ตั้งค่าฟิลก์ Date
	}

	// บันทึก
	if err := entity.DB().Create(&bd).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bd})
}

// GET /deviceborrow/:id
func GetDeviceBorrow(c *gin.Context) {
	var deviceborrow entity.DeviceBorrow
	id := c.Param("id")
	if err := entity.DB().Preload("Member").Preload("DeviceList").Preload("DeviceType").Raw("SELECT * FROM device_borrow WHERE id = ?", id).Find(&deviceborrow).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": deviceborrow})
}

// GET /deviceborrows
func ListDeviceBorrows(c *gin.Context) {
	var deviceborrows []entity.DeviceBorrow
	if err := entity.DB().Preload("Member").Preload("DeviceList").Preload("DeviceType").Raw("SELECT * FROM device_borrows ").Find(&deviceborrows).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": deviceborrows})
}

//function สำหรับลบ deviceborrow ด้วย ID
// DELETE /deviceborrows/:id
func DeleteDeviceBorrow(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM device_borrows WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deviceborrow not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /deviceborrows
func UpdateDeviceBorrow(c *gin.Context) {
	var deviceborrow entity.DeviceBorrow
	if err := c.ShouldBindJSON(&deviceborrow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", deviceborrow.ID).First(&deviceborrow); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deviceborrow not found"})
		return
	}

	if err := entity.DB().Save(&deviceborrow).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": deviceborrow})
}