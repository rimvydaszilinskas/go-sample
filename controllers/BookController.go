package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/bookstore/models"
)

func ListBooks(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		var book []models.Book
		models.DB.Find(&book)
		c.JSON(http.StatusOK, book)
	} else {
		var book models.Book
		models.DB.First(&book, id)

		if book.ID == 0 {
			c.JSON(http.StatusNotFound, book)
			return
		}
		c.JSON(http.StatusOK, book)
	}

}

func AddBook(c *gin.Context) {
	var input CreateBookValidator

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var author models.Author
	models.DB.First(&author, input.Author)

	if author.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		Author:      author,
	}
	models.DB.Create(&book)
	c.JSON(http.StatusCreated, author)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	models.DB.First(&book, id)

	if book.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	var validator UpdateBookValidator
	err1 := c.ShouldBind(&validator)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err1.Error(),
		})
		return
	}

	models.DB.Model(&book).Updates(validator)
	c.JSON(http.StatusOK, book)
}
