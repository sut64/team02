package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Boontita/se-64-example/entity"

)

// POST /book_orders
func CreateBookOrder(c *gin.Context) {

	var bookorder entity.BookOrder
	var booktype entity.BookType
	var company entity.Company
	var orderstatus entity.OrderStatus

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร bookorder
	if err := c.ShouldBindJSON(&bookorder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา book_type ด้วย id
	if tx := entity.DB().Where("id = ?", bookorder.BookTypeID).First(&booktype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booktype not found"})
		return
	}

	// ค้นหา company ด้วย id
	if tx := entity.DB().Where("id = ?", bookorder.CompanyID).First(&company); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company not found"})
		return
	}

	// ค้นหา order_status ด้วย id
	if tx := entity.DB().Where("id = ?", bookorder.OrderStatusID).First(&orderstatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "orderstatus not found"})
		return
	}

	//สร้าง bookorder
	bod := entity.BookOrder{
		BookType: booktype, //โยงความสัมพันธ์กับ Entity BookType
		Company:  company,     //โยงความสัมพันธ์กับ Entity Company
		OrderStatus:  orderstatus,         //โยงความสัมพันธ์กับ Entity OrderStatus
		BookTitle:  bookorder.BookTitle,    //ตั้งค่าฟิลด์ BookTitle
		Author:     bookorder.Author, //ตั้งค่าฟิลด์ Auther
		OrderAmount:  bookorder.OrderAmount, //ตั้งค่าฟิลด์ OrderAmount
		Price: bookorder.Price, //ตั้งค่าฟิลด์ Price
		OrderDate: bookorder.OrderDate, //ตั้งค่าฟิลด์ OrderDate
	}

	// บันทึก
	if err := entity.DB().Create(&bod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bod})
}

// GET /book_order/:id
func GetBookOrder(c *gin.Context) {
	var bookorder entity.BookOrder
	id := c.Param("id")
	if err := entity.DB().Preload("BookType").Preload("Company").Preload("OrderStatus").Raw("SELECT * FROM book_orders WHERE id = ?", id).Find(&bookorder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookorder})
}

// GET /book_orders
func ListBookOrders(c *gin.Context) {
	var bookorders []entity.BookOrder
	if err := entity.DB().Preload("BookType").Preload("Company").Preload("OrderStatus").Raw("SELECT * FROM book_orders ").Find(&bookorders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookorders})
}

//function สำหรับลบ customer ด้วย ID
// DELETE /book_orders/:id
func DeleteBookOrder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM book_orders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookorder not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /book_orders
func UpdateBookOrder(c *gin.Context) {
	var bookorder entity.BookOrder
	if err := c.ShouldBindJSON(&bookorder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bookorder.ID).First(&bookorder); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookorder not found"})
		return
	}

	if err := entity.DB().Save(&bookorder).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookorder})
}
