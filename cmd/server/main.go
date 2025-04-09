package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {

	r := gin.Default()

	r.GET("/image", imagePage)
	r.POST("/create", message)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("server is not started: %v", err)
	}
}

func imagePage(c *gin.Context) {
	filepath := "../../static/image/image.jpg"
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Image not found"})
	}
	c.File(filepath)
}

func message(c *gin.Context) {
	c.String(http.StatusOK, "Bip-bop-bup")
}
