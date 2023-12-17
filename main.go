package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"CRUD/models"
	"CRUD/routes"
)

func main() {
	models.InitDatabase()

	app := fiber.New()

	// Use CORS middleware
    app.Use(cors.New())

	// Routes
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)
	routes.AuthRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

func getBooks(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func createBook(c *fiber.Ctx) error {
	var newBook models.Book
	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	models.DB.Create(&newBook)
	return c.JSON(newBook)
}

func updateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedbook models.Book
	if err := c.BodyParser(&updatedbook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := models.DB.Model(&models.Book{}).Where("id = ?", id).Updates(updatedbook).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.JSON(updatedbook)
}

func deleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := models.DB.Delete(&models.Book{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

