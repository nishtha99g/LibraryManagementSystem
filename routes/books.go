package routes

import (
    "LibraryManagementSystem/controllers"

    "github.com/gofiber/fiber/v2"
)

func BooksRoutes(r *fiber.App) {
	r.Get("/books", controllers.GetBooks)
	r.Get("/books/:id", controllers.GetBook)
	r.Post("/books", controllers.CreateBook)
	r.Put("/books/:id", controllers.UpdateBook)
	r.Delete("/books/:id", controllers.DeleteBook)
}
