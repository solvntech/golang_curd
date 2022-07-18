package services

import (
	"example/demo_crud/src/models"
	"gorm.io/gorm"
)

type IBookService interface {
	CreateBook(book *models.Book) (*models.Book, error)
	GetBook(id string) (*models.Book, error)
	GetAll() *[]models.Book
	UpdateBook(book *models.Book) (*models.Book, error)
	DeleteBook(id string) (*models.Book, error)
}

type BookService struct {
	DB *gorm.DB
}

func NewBookService(DB *gorm.DB) IBookService {
	return &BookService{
		DB,
	}
}

func (bookService *BookService) GetAll() *[]models.Book {
	var books *[]models.Book
	bookService.DB.Find(&books)
	return books
}

func (bookService *BookService) GetBook(id string) (*models.Book, error) {
	var book *models.Book
	db := bookService.DB.Where("id = ?", id).First(&book)
	return book, db.Error
}

func (bookService *BookService) UpdateBook(book *models.Book) (*models.Book, error) {
	if rs := bookService.DB.Where("id = ?", book.ID).First(&book); rs.Error != nil {
		return nil, rs.Error
	}
	bookService.DB.Updates(book)
	return book, nil
}

func (bookService *BookService) DeleteBook(id string) (*models.Book, error) {
	var book *models.Book
	if rs := bookService.DB.Where("id = ?", id).First(&book); rs.Error != nil {
		return nil, rs.Error
	}
	bookService.DB.Delete(book)
	return book, nil
}

func (bookService *BookService) CreateBook(book *models.Book) (*models.Book, error) {
	if rs := bookService.DB.Create(&book); rs.Error != nil {
		return nil, rs.Error
	}
	return book, nil
}
