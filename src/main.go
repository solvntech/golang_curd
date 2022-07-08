package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

type BookForm struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func NewBook(title string, author string, quantity int) *Book {
	book := &Book{
		uuid.New().String(),
		title,
		author,
		quantity,
	}
	return book
}

var books = []Book{
	*NewBook("In Search of Lost Time", "Marcel Proust", 2),
	*NewBook("The Great Gatsby", "F. Scott Fitzgerald", 5),
	*NewBook("War and Peace", "Leo Tolstoy", 6),
}

func findAllBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
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
	var newBookForm BookForm
	if err := ctx.BindJSON(&newBookForm); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "missing data",
		})
		return
	}
	newBook := *NewBook(newBookForm.Title, newBookForm.Author, newBookForm.Quantity)
	books = append(books, newBook)
	ctx.IndentedJSON(http.StatusOK, newBook)
}

func getBookById(id string) (*Book, error) {
	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}

func main() {
	router := gin.Default()
	router.GET("/books", findAllBooks)
	router.GET("/book/:id", findBookById)
	router.POST("/new-book", createBook)
	router.Run(":3000")
}
