package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathan-rr-mello/go-book-author-CRUD/models"
	"github.com/nathan-rr-mello/go-book-author-CRUD/repositories"
)

func (cont *AuthorController) SetupRoutes(route fiber.Router) {
	route.Get("/", cont.FindAllAuthors)
	route.Get("/:id", cont.FindAuthorById)
	route.Post("/", cont.SaveAuthor)
	route.Put("/:id", cont.UpdateAuthor)
	route.Delete("/:id", cont.DeleteAuthor)
}

type AuthorController struct {
	Rep repositories.IAuthorRepository
}

func (cont *AuthorController) FindAllAuthors(c *fiber.Ctx) error {
	//authors, err := repositories.FindAllAuthors()
	authors, err := cont.Rep.FindAllAuthors()
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	return c.JSON(authors)
}

func (cont *AuthorController) FindAuthorById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	//author, err := repositories.FindAuthorById(id)
	author, err := cont.Rep.FindAuthorById(id)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	return c.JSON(author)
}

func (cont *AuthorController) SaveAuthor(c *fiber.Ctx) error {
	var author models.Author
	if err := c.BodyParser(&author); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	//repositories.SaveAuthor(&author)
	cont.Rep.SaveAuthor(&author)
	return c.Status(fiber.StatusCreated).JSON(author)
}

func (cont *AuthorController) UpdateAuthor(c *fiber.Ctx) error {
	var author models.Author
	if err := c.BodyParser(&author); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	//repositories.UpdateAuthor(id, &author)
	cont.Rep.UpdateAuthor(id, &author)
	//savedAuthor, err := repositories.FindAuthorById(id)
	savedAuthor, err := cont.Rep.FindAuthorById(id)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(savedAuthor)
}

func (cont *AuthorController) DeleteAuthor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	//err := repositories.DeleteBook(id)
	err = cont.Rep.DeleteAuthor(id)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).SendString("Sucessfully deleted author...")
}
