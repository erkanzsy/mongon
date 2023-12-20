package main

import (
	"erkanzsy/mongon/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/list-databases", handlers.ListDatabases)
	app.Get("/list-collections/:databaseName", handlers.ListCollections)
	app.Get("/list-documents/:databaseName/:collectionName", handlers.ListDocuments)
	app.Get("/insert-item", handlers.InsertItem)
}
