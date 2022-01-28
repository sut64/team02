package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team02/entity"

)

//function Create เป็นการทำงานแทนคำสั่ง insert ของ SQL
//function นี้จะคืนค่าเป็น devicetype ที่สร้างเสร็จแล้ว กลับไปเป็น JSON ให้ฝั่ง UI นำไปแสดงผล
// POST /devicetypes
func CreateDeviceType(c *gin.Context) {
	var devicetype entity.DeviceType
	if err := c.ShouldBindJSON(&devicetype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&devicetype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": devicetype})
}

//function GetUser จะดึงข้อมูล devicetype ออกมาตาม primary key ที่กำหนด ผ่าน func DB.Raw(...)
// GET /devicetype/:id
func GetDeviceType(c *gin.Context) {
	var devicetype entity.DeviceType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM device_types WHERE id = ?", id).Scan(&devicetype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": devicetype})
}

//ListDeviceTypes จะเป็นการ list รายการของ devicetype ออกมา โดยใช้ array รองรับ
// GET /devicetypes
func ListDeviceTypes(c *gin.Context) {
	var devicetypes []entity.DeviceType
	if err := entity.DB().Raw("SELECT * FROM device_types").Scan(&devicetypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": devicetypes})
}

//function สำหรับลบ devicetype ด้วย ID
// DELETE /devicetypes/:id
func DeleteDeviceType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM device_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "devicetype not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

//function สำหรับ update devicetype ใช้คำสั่ง DB.Save() แทน update ของ SQL
// PATCH /devicetypes
func UpdateDeviceType(c *gin.Context) {
	var devicetype entity.DeviceType
	if err := c.ShouldBindJSON(&devicetype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", devicetype.ID).First(&devicetype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "devicetype not found"})
		return
	}

	if err := entity.DB().Save(&devicetype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": devicetype})
}