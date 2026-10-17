package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Static("/assets", "./assets")
	templates := template.Must(template.ParseGlob("./templates/*.html"))
	server.SetHTMLTemplate(templates)

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
		log.Println("Wrote the name")
	})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}