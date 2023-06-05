package Model

import (
	"Bookstore-Backend-API/pkj/DBConnector"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title     string `gorm:"" json:"title"`
	Author    string `gorm:"" json:"author"`
	Publisher string `gorm:"" json:"publisher"`
}

// automatically called when this package is imported
func init() {
	DBConnector.Connect()
	db = DBConnector.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var resBooks []Book
	db.Find(&resBooks)
	return resBooks
}

func GetBookByID(id int64) Book {
	var bookToBeReturned Book
	db.First(&bookToBeReturned, id)
	return bookToBeReturned
}

func CreateBook(bookToBeCreated *Book) *Book {
	db.Create(&bookToBeCreated)
	return bookToBeCreated
}

func DeleteBook(id int64) {
	var bookToBeDeleted Book
	db.First(&bookToBeDeleted, id)
	db.Delete(&bookToBeDeleted)
}

func UpdateBook(id int64, newBook Book) Book {
	oldBook := GetBookByID(id)
	if newBook.Title != "" {
		oldBook.Title = newBook.Title
	}
	if newBook.Author != "" {
		oldBook.Author = newBook.Author
	}
	if newBook.Publisher != "" {
		oldBook.Publisher = newBook.Publisher
	}

	//now the oldBook has been updated

	db.Save(&oldBook)
	return oldBook
}
