package routes

import (
    "CRUD/controllers"

    "github.com/gofiber/fiber/v2"
)

func AuthRoutes(r *fiber.App) {
    r.Post("/login", controllers.Login)
    r.Post("/signup", controllers.Signup)
    r.Get("/home", controllers.Home)
    r.Get("/logout", controllers.Logout)
}
