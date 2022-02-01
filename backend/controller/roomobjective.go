package controller

import (
	"net/http"

	"github.com/sut64/team02/entity"
	"github.com/gin-gonic/gin"
)

// POST /roomobjective
func CreateRoomObjective(c *gin.Context) {
	var roomobjective entity.RoomObjective
	if err := c.ShouldBindJSON(&roomobjective); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&roomobjective).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomobjective})
}

// GET /room_objective/:id
func GetRoomObjective(c *gin.Context) {
	var roomobjective entity.RoomObjective
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM room_objectives WHERE id = ?", id).Scan(&roomobjective).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomobjective})
}

// GET /room_objective
func ListRoomObjective(c *gin.Context) {
	var roomobjective []entity.RoomObjective
	if err := entity.DB().Raw("SELECT * FROM room_objectives").Scan(&roomobjective).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": roomobjective})
}

// DELETE /room_objectives/:id
func DeleteRoomObjective(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_objectives WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomobjective not found"})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"data": id})
}