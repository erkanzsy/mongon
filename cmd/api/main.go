package main

import (
	"erkanzsy/mongon/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, World!")

	database.ConnectMongo()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
