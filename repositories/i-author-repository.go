package repositories

import "github.com/nathan-rr-mello/go-book-author-CRUD/models"

type IAuthorRepository interface {
	FindAllAuthors() ([]models.Author, error)
	FindAuthorById(id int) (models.Author, error)
	SaveAuthor(author *models.Author)
	UpdateAuthor(id int, author *models.Author)
	DeleteAuthor(id int) error
}
