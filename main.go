package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// bytes, err := os.ReadFile("./index.html")

	server := gin.Default()
	server.Static("/assets", "./assets")
	// server.LoadHTMLGlob("templates/*")
	html := template.Must(template.ParseFiles("index.html"))

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H {
			"title": "Hello World",
		})
	})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
