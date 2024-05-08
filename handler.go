package main

import (
	"github.com/gin-gonic/gin"
)

func createShortURL(c *gin.Context) {
	// Handle short url creation
	// 1. Convert the URL to the some short code using short code generator that in unique
	// 2. Persist the shortCode with original URL mapping for the faster lookups
	// 3. Return the shortCode URL to the user as API response
}

func handleShortURL(c *gin.Context) {
	// Handle the short URLs here.
	// 1. Resolve the original URL using the shortPath string
	// 2. If the original URL is found, redirect to original URL
	// 3. Throw default 404 Not found if the original Path is not found.
}

func main() {
	router := gin.Default()

	router.POST("/api/v1/shorturl/create", createShortURL)
	router.Any("/:shortPath", handleShortURL)
	router.Run(":8080")
}
