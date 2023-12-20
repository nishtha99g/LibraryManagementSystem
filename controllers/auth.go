package controllers

import (
    "LibraryManagementSystem/models"
    "time"

    "LibraryManagementSystem/utils"

    "github.com/dgrijalva/jwt-go"
    "github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("my_secret_key")

func Login(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var existingUser models.User

    models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User does not exist"})
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Password"})
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not generate token"})
	}

	// Set JWT token as a cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HTTPOnly: true,
		SameSite: "Strict",
		Secure:   true,
	})

	return c.JSON(fiber.Map{"success": "user logged in"})
}

func Signup(c *fiber.Ctx) error {
	
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var existingUser models.User
	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}

	hashedPassword, err := utils.GenerateHashPassword(user.Password)


	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not generate hash password"})
	}

	user.Password = hashedPassword

	models.DB.Create(&user)

	return c.JSON(user)
}

func Home(c *fiber.Ctx) error {
	cookie := c.Cookies("token")

	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if claims.Role != "user" && claims.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.JSON(fiber.Map{"success": "home page", "role": claims.Role})
}

func Logout(c *fiber.Ctx) error {
	// Set an expired token by clearing the existing token cookie
	c.ClearCookie("token")

	// Respond with a JSON success message
	return c.JSON(fiber.Map{"success": "user logged out"})
}
