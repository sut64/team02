package controller
 
import (
	"github.com/sut64/team02/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/asaskevich/govalidator"
)

// POST /BookingRoom
func CreateBookingRoom(c *gin.Context) {

	var bookingroom entity.BookingRoom
	var roomandtime entity.RoomAndTime
	var roomtype entity.RoomType
	var roomobjective entity.RoomObjective

	if err := c.ShouldBindJSON(&bookingroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bookingroom.RoomAndTimeID).First(&roomandtime); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomandtime not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", bookingroom.RoomTypeID).First(&roomtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomtype not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", bookingroom.RoomObjectiveID).First(&roomobjective); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomobjective not found"})
		return
	}


// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(bookingroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	booking := entity.BookingRoom{
		RoomAndTime: 		roomandtime, 					
		RoomType:      		roomtype,      				
		RoomObjective:      roomobjective,          				
		PhoneBooker:       	bookingroom.PhoneBooker, 			
		QuantityMember:		bookingroom.QuantityMember,
		BookingRoomAt:		bookingroom.BookingRoomAt, 
	}

	if err := entity.DB().Create(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// GET /bookingroom/:id
func GetBookingroom(c *gin.Context) {
	var bookingroom entity.BookingRoom
	id := c.Param("id")
	if err := entity.DB().Preload("RoomAndTime").Preload("RoomType").Preload("RoomObjective").Raw("SELECT * FROM booking_rooms WHERE id = ?", id).Find(&bookingroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookingroom})
}

// GET /bookingrooms
func ListBookingrooms(c *gin.Context) {
	var bookingrooms []entity.BookingRoom
	if err := entity.DB().Preload("RoomAndTime").Preload("RoomType").Preload("RoomObjective").Raw("SELECT * FROM booking_rooms").Find(&bookingrooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookingrooms})
}

// DELETE /bookingrooms/:id
func DeleteBookingRoom(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM booking_rooms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookingroom not found"})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /bookingrooms
func UpdateBookingRoom(c *gin.Context) {
	var bookingroom entity.BookingRoom
	if err := c.ShouldBindJSON(&bookingroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	if tx := entity.DB().Where("id = ?", bookingroom.ID).First(&bookingroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookingroom not found"})
		return
	}
 
	if err := entity.DB().Save(&bookingroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"data": bookingroom})
}