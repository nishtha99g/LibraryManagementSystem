package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"LibraryManagementSystem/models"
	"LibraryManagementSystem/routes"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"os"
	"LibraryManagementSystem/middlewares"
	"LibraryManagementSystem/controllers"
)

// Redis client
var redisClient *redis.Client

// Logrus logger instance
var logger = logrus.New()

func init() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// Set logrus formatter to JSON format
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Set logrus output to a file
	file, err := os.OpenFile("logs.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
}

func main() {
	models.InitDatabase()

	app := fiber.New()

	// Use CORS middleware
    app.Use(cors.New())

	app.Use(middlewares.PrometheusMiddleware)

	controllers.SetRedisClient(redisClient)

	// Routes
	routes.AuthRoutes(app)
	routes.BooksRoutes(app)

	//app.Get("/metrics", middlewares.PrometheusHandler())

	log.Fatal(app.Listen(":8080"))
}



