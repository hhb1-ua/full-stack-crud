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

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("dist/index.html")
	})

	api := router.Group("/api")
	{
		api.GET("/authors", GetAuthors)
		api.GET("/authors/:id", GetAuthorByID)
		api.POST("/authors", CreateAuthor)
		api.PUT("/authors/:id", UpdateAuthor)
		api.DELETE("/authors/:id", DeleteAuthor)

		api.GET("/books/:id", GetBooksByAuthorID)
		api.POST("/books", CreateBook)
		api.PUT("/books/:id", UpdateBook)
		api.DELETE("/books/:id", DeleteBook)
	}

	router.Run(":8080")
}
