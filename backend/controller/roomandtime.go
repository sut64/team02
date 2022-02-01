package controller

import (
	"net/http"

	"github.com/sut64/team02/entity"
	"github.com/gin-gonic/gin"
)

// POST /roomandtime
func CreateRoomAndTime(c *gin.Context) {
	var roomandtime entity.RoomAndTime
	if err := c.ShouldBindJSON(&roomandtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&roomandtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomandtime})
}

// GET /room_and_time/:id
func GetRoomAndTime(c *gin.Context) {
	var roomandtime entity.RoomAndTime
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM room_and_times WHERE id = ?", id).Scan(&roomandtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomandtime})
}

// GET /room_and_time
func ListRoomAndTime(c *gin.Context) {
	var roomandtime []entity.RoomAndTime
	if err := entity.DB().Raw("SELECT * FROM room_and_times").Scan(&roomandtime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomandtime})
}

// DELETE /roomandtime/:id
func DeleteRoomAndTime(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_and_times WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomandtime not found"})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"data": id})
}