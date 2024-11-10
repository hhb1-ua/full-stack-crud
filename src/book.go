package main

import (
	"time"

	"gorm.io/gorm"
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
	gorm.Model
	Title       string
	Description *string
	Category    Category
	Publication time.Time
	AuthorID    uint
}
