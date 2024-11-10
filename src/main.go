package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	{
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database")
		}
		db.AutoMigrate(&Author{}, &Book{})
	}

	router := gin.Default()

	router.GET("/api/authors", GetAuthors)
	router.GET("/api/authors/:id", GetAuthorByID)
	router.POST("/api/authors", PostAuthor)
	router.PUT("/api/authors/:id", PutAuthor)
	router.DELETE("/api/authors/:id", DeleteAuthor)

	router.Run(":8080")
}
