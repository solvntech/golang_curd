package controllers

import (
	"example/demo_crud/src/helper"
	"example/demo_crud/src/models"
	"example/demo_crud/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IBookController interface {
	FindAll(context *gin.Context)
	FindById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type BookController struct {
	bookService services.IBookService
}

func NewBookController(bookService services.IBookService) IBookController {
	return &BookController{
		bookService,
	}
}

func (bookController BookController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, bookController.bookService.GetAll())
}

func (bookController BookController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := bookController.bookService.GetBook(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, book)
}

func (bookController BookController) Create(ctx *gin.Context) {
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
	book, err := bookController.bookService.CreateBook(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, book)
}

func (bookController BookController) Update(ctx *gin.Context) {
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
	book, err := bookController.bookService.UpdateBook(&newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, book)
}

func (bookController BookController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := bookController.bookService.DeleteBook(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.IndentedJSON(http.StatusOK, book)
}
