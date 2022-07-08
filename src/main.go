package main

import (
	"errors"
	"example/demo_crud/src/database"
	"example/demo_crud/src/helper"
	"example/demo_crud/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func findAllBooks(ctx *gin.Context) {
	var books []models.Book
	database.DBInstance.Find(&books)
	ctx.JSON(http.StatusOK, &books)
}

func findBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := getBookById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, book)
}

func createBook(ctx *gin.Context) {
	var newBook models.Book
	if err := ctx.BindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing data",
		})
		return
	}
	if ok, err := helper.Validate(newBook); !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if rs := database.DBInstance.Create(&newBook); rs.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": rs.Error.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, newBook)
}

func getBookById(id string) (*models.Book, error) {
	//for _, book := range books {
	//	if book.ID == id {
	//		return &book, nil
	//	}
	//}
	return nil, errors.New("book not found")
}

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// connect to db
	db := database.ConnectDB()
	fmt.Println(db)

	router := gin.Default()
	router.GET("/initDB", func(ctx *gin.Context) {
		database.MigrateDB()
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	router.GET("/books", findAllBooks)
	router.GET("/book/:id", findBookById)
	router.POST("/new-book", createBook)
	router.Run(":3000")
}
