package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/bookstore/controllers"
	"github.com/rimvydaszilinskas/bookstore/models"
)

func main() {
	models.Connect()
	defer models.DB.Close()
	models.Sync()

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"name": "name",
		})
	})

	app.GET("/authors", controllers.ListAuthors)
	app.POST("/authors", controllers.AddAuthor)
	app.PATCH("/authors/:id", controllers.UpdateAuthor)

	app.GET("/books", controllers.ListBooks)
	app.GET("/books/:id", controllers.ListBooks)
	app.POST("/books", controllers.AddBook)
	app.PATCH("/books/:id", controllers.UpdateBook)

	app.Run(":8080")
}