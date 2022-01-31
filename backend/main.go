package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/controller"
	"github.com/sut64/team02/entity"
	"github.com/sut64/team02/middlewares"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// BookInformations Routes
			protected.GET("/book_informations", controller.ListBookInformations)
			protected.GET("/bookinformation/:id", controller.GetBookInformation)
			protected.GET("/bookinformation/bookType/:id", controller.GetInfoByID)
			protected.POST("/book_informations", controller.CreateBookInformation)
			protected.PATCH("/book_informations", controller.UpdateBookInformation)
			protected.DELETE("/book_informations/:id", controller.DeleteBookInformation)

			// BookType Routes
			protected.GET("/book_types", controller.ListBookTypes)
			protected.GET("/booktype/:id", controller.GetBookType)
			protected.POST("/book_types", controller.CreateBookType)
			protected.PATCH("/book_types", controller.UpdateBookType)
			protected.DELETE("/book_types/:id", controller.DeleteBookType)

			// BookLocation Routes
			protected.GET("/book_locations", controller.ListBookLocations)
			protected.GET("/booklocation/:id", controller.GetBookLocation)
			protected.POST("/book_locations", controller.CreateBookLocation)
			protected.PATCH("/book_locations", controller.UpdateBookLocation)
			protected.DELETE("/book_locations/:id", controller.DeleteBookLocation)

			// Company Routes
			protected.GET("/companies", controller.ListCompanies)
			protected.GET("/company/:id", controller.GetCompany)
			protected.POST("/compamies", controller.CreateCompany)
			protected.PATCH("/companies", controller.UpdateCompany)
			protected.DELETE("/companies/:id", controller.DeleteCompany)

			//OrderStatus Routes
			protected.GET("/order_statuses", controller.ListOrderStatuses)
			protected.GET("/order_status/:id", controller.GetOrderStatus)
			protected.POST("/order_statuses", controller.CreateOrderStatus)
			protected.PATCH("/order_statuses", controller.UpdateOrderStatus)
			protected.DELETE("/order_statuses/:id", controller.DeleteOrderStatus)

			// members Routes
			protected.GET("/members", controller.ListMembers)
			protected.GET("/member/:id", controller.GetMember)
			protected.PATCH("/members", controller.UpdateMember)
			protected.DELETE("/members/:id", controller.DeleteMember)

			// BookOrder Routes
			protected.GET("/book_orders", controller.ListBookOrders)
			protected.GET("/book_order/:id", controller.GetBookOrder)
			protected.POST("/book_orders", controller.CreateBookOrder)
			protected.PATCH("/book_orders", controller.UpdateBookOrder)
			protected.DELETE("/book_orders/:id", controller.DeleteBookOrder)

			// Borrow Routes
			protected.GET("/borrows", controller.ListBorrows)
			protected.GET("/borrow/:id", controller.GetBorrow)
			protected.GET("/borrow/member/:id", controller.ListBorrowsByMember)
			protected.POST("/borrows", controller.CreateBorrow)
			protected.PATCH("/borrows", controller.UpdateBorrow)
			protected.DELETE("/borrows/:id", controller.DeleteBorrow)

			// Status Routes
			protected.GET("/statuses", controller.ListStatuses)
			protected.GET("/status/:id", controller.GetStatus)
			protected.POST("/statuses", controller.CreateStatus)
			protected.PATCH("/statuses", controller.UpdateStatus)
			protected.DELETE("/statuses/:id", controller.DeleteStatus)

			// ServicePlace Routes
			protected.GET("/places", controller.ListPlaces)
			protected.GET("/place/:id", controller.GetPlace)
			protected.POST("/places", controller.CreatePlace)
			protected.PATCH("/places", controller.UpdatePlace)
			protected.DELETE("/places/:id", controller.DeletePlace)

			// DeviceList Routes
			protected.GET("/devicelists", controller.ListDeviceLists)
			protected.GET("/devicelist/:id", controller.GetDeviceList)
			protected.POST("/devicelists", controller.CreateDeviceList)
			protected.PATCH("/devicelists", controller.UpdateDeviceList)
			protected.DELETE("/devicelists/:id", controller.DeleteDeviceList)
	  
			// DeviceType Routes
			protected.GET("/devicetypes", controller.ListDeviceTypes)
			protected.GET("/devicetype/:id", controller.GetDeviceType)
			protected.POST("/devicetypes", controller.CreateDeviceType)
			protected.PATCH("/devicetypes", controller.UpdateDeviceType)
			protected.DELETE("/devicetypes/:id", controller.DeleteDeviceType)
	  
			// DeviceBorrow Routes
			protected.GET("/deviceborrows", controller.ListDeviceBorrows)
			protected.GET("/deviceborrow/:id", controller.GetDeviceBorrow)
			protected.POST("/deviceborrows", controller.CreateDeviceBorrow)
			protected.PATCH("/deviceborrows", controller.UpdateDeviceBorrow)
			protected.DELETE("/deviceborrows/:id", controller.DeleteDeviceBorrow)
			
		}
	}

	// members Routes
	r.POST("/members", controller.CreateMember)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
