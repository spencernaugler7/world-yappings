package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Static("/assets", "./assets")
	server.LoadHTMLGlob("templates/*")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello World",
			"name":  "Ben Tennison",
		})
		log.Println("Wrote the name")
	})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
