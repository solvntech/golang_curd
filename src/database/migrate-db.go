package database

import "example/demo_crud/src/models"

func MigrateDB() {
	books := []models.Book{
		{Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 6},
		{Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 5},
	}
	DBInstance.Migrator().DropTable(&models.Book{})
	DBInstance.AutoMigrate(&models.Book{})
	DBInstance.Create(&books)
}
