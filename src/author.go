package main

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name     string
	Email    string
	Birthday *time.Time
	Books    []Book
}
