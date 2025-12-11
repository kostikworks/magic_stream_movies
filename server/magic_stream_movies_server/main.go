package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kostikworks/magic_stream_movies/server/magic_stream_movies_server/database"
	"github.com/kostikworks/magic_stream_movies/server/magic_stream_movies_server/routes"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	//this is the main function

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Magic Stream Movies!")
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: failed to load .env file")
	}

	var client *mongo.Client = database.ConnectDB()

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to reach server: %v", err)
	}

	defer func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	routes.SetupUnprotectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
