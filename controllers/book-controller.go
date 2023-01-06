package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathan-rr-mello/go-book-author-CRUD/models"
	"github.com/nathan-rr-mello/go-book-author-CRUD/repositories"
)

func (cont *BookController) SetupRoutes(route fiber.Router) {
	route.Get("/", cont.FindAllBooks)
	route.Get("/:id", cont.FindBookById)
	route.Post("/", cont.SaveBook)
	route.Put("/:id", cont.UpdateBook)
	route.Delete(":id", cont.DeleteBook)
}

type BookController struct {
	Rep repositories.IBookRepository
}

func (cont *BookController) FindAllBooks(c *fiber.Ctx) error {
	books := cont.Rep.FindAllBooks()
	return c.Status(fiber.StatusOK).JSON(books)
}

func (cont *BookController) FindBookById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	book, err := cont.Rep.FindBookById(id)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func (cont *BookController) SaveBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	cont.Rep.SaveBook(&book)
	return c.Status(fiber.StatusCreated).JSON(book)
}

func (cont *BookController) UpdateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	cont.Rep.UpdateBook(id, &book)
	savedBook, err := cont.Rep.FindBookById(id)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(savedBook)
}

func (cont *BookController) DeleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	if err := cont.Rep.DeleteBook(id); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).SendString("Sucessfully deleted book...")
}
