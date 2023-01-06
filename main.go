package main

import (
	"log"
	"github.com/nathan-rr-mello/go-book-author-CRUD/repositories"
	"github.com/nathan-rr-mello/go-book-author-CRUD/database"
	"github.com/nathan-rr-mello/go-book-author-CRUD/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	bookRoute := app.Group("/book")
	authorRoute := app.Group("/author")
	bookController := &controllers.BookController{
		Rep: &repositories.BookRepository{},
	}
	bookController.SetupRoutes(bookRoute)
	authorController := &controllers.AuthorController{
		Rep: &repositories.AuthorRepository{},
	}
	authorController.SetupRoutes(authorRoute)
}

func main() {
	database.ConnectDB()
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
