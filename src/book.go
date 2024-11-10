package main

import (
	"time"
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
