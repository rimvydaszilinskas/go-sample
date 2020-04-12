package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rimvydaszilinskas/bookstore/models"
	"net/http"
)

func ListAuthors(c *gin.Context) {
	var authors []models.Author
	models.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func AddAuthor(c *gin.Context) {
	var input CreateAuthorValidator

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	author := models.Author{
		Name: input.Name,
		Introduction: input.Introduction,
	}
	models.DB.Create(&author)

	c.JSON(http.StatusCreated, author)
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	models.DB.First(&author, id)

	if author.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	var validator UpdateAuthorValidator
	err1 := c.ShouldBind(&validator)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err1.Error(),
		})
		return
	}

	models.DB.Model(&author).Updates(validator)

	c.JSON(http.StatusOK, author)
}