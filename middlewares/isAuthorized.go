// PATH: go-auth/middlewares/isAuthorized.go

package middlewares

import (
    "LibraryManagementSystem/utils"

    "github.com/gofiber/fiber/v2"
)

func IsAuthorized() fiber.Handler {
    return func(c *fiber.Ctx) error {
        cookie, err := c.Cookie("token")

        if err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
        }

        claims, err := utils.ParseToken(cookie)

        if err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
        }

        c.Locals("role", claims.Role)
        return c.Next()
    }
}
