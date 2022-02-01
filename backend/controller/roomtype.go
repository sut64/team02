package controller

import (
	"net/http"

	"github.com/sut64/team02/entity"
	"github.com/gin-gonic/gin"
)

// POST /roomtype
func CreateRoomType(c *gin.Context) {
	var roomtype entity.RoomType
	if err := c.ShouldBindJSON(&roomtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// GET /room_types/:id
func GetRoomType(c *gin.Context) {
	var roomtype entity.RoomType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM room_types WHERE id = ?", id).Scan(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// GET /room_types
func ListRoomtype(c *gin.Context) {
	var roomtype []entity.RoomType
	if err := entity.DB().Raw("SELECT * FROM room_types").Scan(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// DELETE /roomtypes/:id
func DeleteRoomType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomtype not found"})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"data": id})
}