package repositories

import (
	"fmt"

	"github.com/nathan-rr-mello/go-book-author-CRUD/database"
	"github.com/nathan-rr-mello/go-book-author-CRUD/models"
)

type BookRepository struct{}

func (rep *BookRepository) FindAllBooks() []models.Book {
	books := []models.Book{}
	database.Db.Find(&books)
	return books
}

func (rep *BookRepository) FindBookById(id int) (models.Book, error) {
	book := models.Book{}
	database.Db.Find(&book, "id = ?", id)
	if book.ID == 0 {
		return models.Book{}, fmt.Errorf("book with id %v does not exist", id)
	}
	return book, nil
}

func (rep *BookRepository) SaveBook(book *models.Book) {
	database.Db.Create(book)
}

func (rep *BookRepository) UpdateBook(id int, book *models.Book) {
	database.Db.Model(models.Book{}).Where("id = ?", id).Updates(book)
}

func (rep *BookRepository) DeleteBook(id int) error {
	return database.Db.Delete(&models.Book{}, id).Error
}
