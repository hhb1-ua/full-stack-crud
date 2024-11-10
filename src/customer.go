package main

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name     string
	Email    string
	Birthday *time.Time
}
