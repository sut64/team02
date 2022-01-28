package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sut64/team02/controller"
	"github.com/sut64/team02/se-64/entity"
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


			// BookOrders Routes
			protected.GET("/book_orders", controller.ListBookOrders)
			protected.GET("/bookorder/:id", controller.GetBookOrder)
			protected.POST("/book_orders", controller.CreateBookOrder)
			protected.PATCH("/book_orders", controller.UpdateBookOrder)
			protected.DELETE("/book_orders/:id", controller.DeleteBookOrder)

			// members Routes
			protected.GET("/members", controller.ListMembers)
			protected.GET("/member/:id", controller.GetMember)
			
			protected.PATCH("/members", controller.UpdateMember)
			protected.DELETE("/members/:id", controller.DeleteMember)

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
