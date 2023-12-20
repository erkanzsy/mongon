package handlers

import (
	"context"
	"erkanzsy/mongon/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func ListDatabases(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	databases, err := database.MClient.Client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(databases)
}

func ListCollections(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Get the database name from the route parameters
	databaseName := c.Params("databaseName")

	// Specify the database
	database := database.MClient.Client.Database(databaseName)

	// Get list of collections in the specified database
	collections, err := database.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Return the list of collections as JSON
	return c.JSON(collections)
}

func ListDocuments(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Get the database and collection names from the route parameters
	databaseName := c.Params("databaseName")
	collectionName := c.Params("collectionName")

	// Specify the database and collection
	database := database.MClient.Client.Database(databaseName)
	collection := database.Collection(collectionName)

	// Find all documents in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	defer cursor.Close(ctx)

	// Decode the documents into a slice
	var documents []bson.M
	if err := cursor.All(ctx, &documents); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Return the list of documents as JSON
	return c.JSON(documents)
}

func InsertItem(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := bson.M{
		"name":  "John Doe",
		"email": "john.doe@example.com",
		"age":   30,
	}

	// Specify the database and collection
	database := database.MClient.Client.Database("mongon")
	collection := database.Collection("mongon")

	// Insert the data into the collection
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Return the ID of the inserted document
	return c.JSON(result.InsertedID)
}
