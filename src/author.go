package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Author struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Birthday *time.Time `json:"birthday"`
	Books    []Book     `json:"books"`
}

func GetAuthors(c *gin.Context) {}

func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")

	var author Author
	db.First(&author, id)

}

func PostAuthor(c *gin.Context) {}

func PutAuthor(c *gin.Context) {}

func DeleteAuthor(c *gin.Context) {}
