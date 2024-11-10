package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Category string

const (
	Unknown  Category = "Unknown"
	Fantasy  Category = "Fantasy"
	Fiction  Category = "Fiction"
	Romance  Category = "Romance"
	Horror   Category = "Horror"
	Thriller Category = "Thriller"
	Mystery  Category = "Mystery"
)

type Book struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Category    Category  `json:"category" binding:"required"`
	Publication time.Time `json:"publication" binding:"required"`
	AuthorID    uint      `json:"author" binding:"required"`
}

func GetBooksByAuthorID(c *gin.Context) {
	var books []Book

	if err := db.Find(&books).Where("AuthorID == ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't fetch from author"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book Book

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var book Book

	if err := db.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	var book Book

	if err := db.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
