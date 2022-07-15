package controller

import (
	"strconv"

	"github.com/Wmartevi/web-api-go/database"
	"github.com/Wmartevi/web-api-go/model"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book model.Book
	err = db.First(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)

}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()
	var books model.Book

	err := c.ShouldBindJSON(&books)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}
	c.JSON(200, books)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()
	var books []model.Book
	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list book: " + err.Error(),
		})
		return
	}
	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()
	var book model.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book model.Book
	err = db.Delete(&book, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}
	c.Status(204)

}
