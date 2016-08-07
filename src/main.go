package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("APP_ENVIRONMENT") == "dev" {
		fmt.Println("\n****************** APP READY\n")
	}

	r := gin.Default()
	r.GET("/", onRoot)

	r.Run(":" + os.Getenv("HTTP_PORT"))
}

func onRoot(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
