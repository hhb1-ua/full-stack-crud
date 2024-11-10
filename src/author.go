package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Author struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name" binding:"required"`
	Email    string     `json:"email" binding:"required,email"`
	Birthday *time.Time `json:"birthday" binding:"required"`
	Books    []Book     `json:"books" binding:"required"`
}

func GetAuthors(c *gin.Context) {
	var authors []Author

	if err := db.Find(&authors).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't fetch the authors"})
		return
	}

	c.JSON(http.StatusOK, authors)
}

func GetAuthorByID(c *gin.Context) {
	var author Author

	if err := db.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func CreateAuthor(c *gin.Context) {
	var author Author

	if err := c.BindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create author"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func UpdateAuthor(c *gin.Context) {
	var author Author

	if err := db.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	if err := c.BindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update author"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c *gin.Context) {
	var author Author

	if err := db.First(&author, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	if err := db.Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
