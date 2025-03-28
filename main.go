package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("X-Timestamp", time.Now().Format(time.RFC3339))
		c.Next()
	})

	// API エンドポイント
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	router.POST("/api", func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			return
		}

		bodyString := string(bodyBytes)

		fmt.Println("Received POST request body:")
		fmt.Println(bodyString)

		c.JSON(http.StatusOK, gin.H{
			"message":  "Request body logged to terminal",
			"received": bodyString,
		})
	})

	router.StaticFile("/", "./public/index.html")
	router.Static("/static", "./public")

	router.Run(":3000")
}
