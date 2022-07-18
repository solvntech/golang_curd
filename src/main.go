package main

import (
	"example/demo_crud/src/controllers"
	"example/demo_crud/src/database"
	"example/demo_crud/src/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
)

var (
	DB             *gorm.DB
	bookService    services.IBookService
	bookController controllers.IBookController
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// connect to db
	DB = database.ConnectDB()
	bookService = services.NewBookService(DB)
	bookController = controllers.NewBookController(bookService)

	router := gin.Default()
	router.GET("/initDB", func(ctx *gin.Context) {
		database.MigrateDB()
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	router.GET("/books", bookController.FindAll)
	router.GET("/book/:id", bookController.FindById)
	router.POST("/new-book", bookController.Create)
	router.Run(":3000")
}
