package controllers

import (
	"github.com/gofiber/fiber/v2"
	"LibraryManagementSystem/models"
	"github.com/go-redis/redis/v8"
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/sirupsen/logrus"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client // Declare redisClient at the package level
	logger      = logrus.New()
)

// Function to set the global redis client
func SetRedisClient(client *redis.Client) {
	redisClient = client
}

func GetBooks(c *fiber.Ctx) error {
	logger.Info("Handling GET /books request")

	// Define cache key
	cacheKey := "books"

	// Try to get data from cache
	cacheData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Cache hit, parse JSON and return
		var books []models.Book
		if err := json.Unmarshal([]byte(cacheData), &books); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse cache data"})
		}
		fmt.Println("Cached data:", books)
		return c.JSON(books)
	}

	// Cache miss, fetch data from the database
	fmt.Println("Fetching data from DB")
	var books []models.Book
	models.DB.Find(&books)

	// Marshal data to JSON
	booksJSON, err := json.Marshal(books)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to marshal data"})
	}

	// Set data to cache with expiration (1 hour)
	err = redisClient.SetEX(ctx, cacheKey, booksJSON, time.Hour).Err()
	if err != nil {
		fmt.Println("Failed to set data to cache:", err)
	}

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var newBook models.Book
	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	models.DB.Create(&newBook)

	logger.WithFields(logrus.Fields{
		"book_id": newBook.ID,
	}).Info("Book created successfully")

	return c.JSON(newBook)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedbook models.Book
	if err := c.BodyParser(&updatedbook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if err := models.DB.Model(&models.Book{}).Where("id = ?", id).Updates(updatedbook).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	// Invalidate the cache after update
    cacheKey := "books"
    if err := redisClient.Del(ctx, cacheKey).Err(); err != nil {
        fmt.Println("Failed to delete cache:", err)
    }

	logger.WithFields(logrus.Fields{
		"book_id": id,
	}).Info("Book updated successfully")

	return c.JSON(updatedbook)
}

func DeleteBook(c *fiber.Ctx) error {
	logger.WithFields(logrus.Fields{
		"book_id": c.Params("id"),
	}).Info("Handling DELETE /books/:id request")
	id := c.Params("id")
	if err := models.DB.Delete(&models.Book{}, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	// Invalidate the cache after update
    cacheKey := "books"
    if err := redisClient.Del(ctx, cacheKey).Err(); err != nil {
        fmt.Println("Failed to delete cache:", err)
    }

	return c.SendStatus(fiber.StatusNoContent)
}