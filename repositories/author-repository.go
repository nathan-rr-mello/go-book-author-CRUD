package repositories

import (
	"github.com/nathan-rr-mello/go-book-author-CRUD/database"
	"github.com/nathan-rr-mello/go-book-author-CRUD/models"

	"fmt"
)

type AuthorRepository struct{}

func (rep *AuthorRepository) FindAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	err := database.Db.Model(&models.Author{}).Preload("Books").Find(&authors).Error
	return authors, err
}

func (rep *AuthorRepository) FindAuthorById(id int) (models.Author, error) {
	author := models.Author{}
	database.Db.Model(&models.Author{}).Preload("Books").Find(&author, "id = ?", id)
	if author.ID == 0 {
		return models.Author{}, fmt.Errorf("author with id %v does not exist", id)
	}
	return author, nil
}

func (rep *AuthorRepository) SaveAuthor(author *models.Author) {
	database.Db.Create(author)
}

func (rep *AuthorRepository) UpdateAuthor(id int, author *models.Author) {
	database.Db.Model(models.Author{}).Where("id = ?", id).Updates(author)
}

func (rep *AuthorRepository) DeleteAuthor(id int) error {
	return database.Db.Delete(&models.Author{}, id).Error
}
