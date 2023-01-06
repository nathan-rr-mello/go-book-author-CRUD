package repositories

import "github.com/nathan-rr-mello/go-book-author-CRUD/models"

type IBookRepository interface {
	FindAllBooks() []models.Book
	FindBookById(id int) (models.Book, error)
	SaveBook(author *models.Book)
	UpdateBook(id int, book *models.Book)
	DeleteBook(id int) error
}
