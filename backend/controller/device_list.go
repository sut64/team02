package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team02/entity"

	
)

//function Create เป็นการทำงานแทนคำสั่ง insert ของ SQL
//function นี้จะคืนค่าเป็น devicelist ที่สร้างเสร็จแล้ว กลับไปเป็น JSON ให้ฝั่ง UI นำไปแสดงผล
// POST /devicelists
func CreateDeviceList(c *gin.Context) {
	var devicelist entity.DeviceList
	if err := c.ShouldBindJSON(&devicelist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&devicelist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": devicelist})
}

//function GetUser จะดึงข้อมูล devicelist ออกมาตาม primary key ที่กำหนด ผ่าน func DB.Raw(...)
// GET /devicelist/:id
func GetDeviceList(c *gin.Context) {
	var devicelist entity.DeviceList
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM device_lists WHERE id = ?", id).Scan(&devicelist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": devicelist})
}

//ListDeviceLists จะเป็นการ list รายการของ devicelist ออกมา โดยใช้ array รองรับ
// GET /devicelists
func ListDeviceLists(c *gin.Context) {
	var devicelists []entity.DeviceList
	if err := entity.DB().Raw("SELECT * FROM device_lists").Scan(&devicelists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": devicelists})
}

//function สำหรับลบ devicelist ด้วย ID
// DELETE /devicelists/:id
func DeleteDeviceList(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM device_lists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "devicelist not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

//function สำหรับ update devicelist ใช้คำสั่ง DB.Save() แทน update ของ SQL
// PATCH /devicelists
func UpdateDeviceList(c *gin.Context) {
	var devicelist entity.DeviceList
	if err := c.ShouldBindJSON(&devicelist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", devicelist.ID).First(&devicelist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "devicelist not found"})
		return
	}

	if err := entity.DB().Save(&devicelist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": devicelist})
}