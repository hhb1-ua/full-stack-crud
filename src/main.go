package main

import (
	"fmt"
	"net/http"
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

	router.StaticFS("/static", http.Dir("dist"))

	router.GET("/", func(ctx *gin.Context) {
		ctx.File("dist/index.html")
	})

	api := router.Group("/api")
	{
		api.GET("api/authors", GetAuthors)
		api.GET("api/authors/:id", GetAuthorByID)
		api.POST("api/authors", CreateAuthor)
		api.PUT("api/authors/:id", UpdateAuthor)
		api.DELETE("api/authors/:id", DeleteAuthor)

		api.GET("api/books/:id", GetBooksByAuthorID)
		api.POST("api/books", CreateBook)
		api.PUT("api/books/:id", UpdateBook)
		api.DELETE("api/books/:id", DeleteBook)
	}

	router.Run(":8080")
}
