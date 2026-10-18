package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"world.com/yappings/data"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err := data.SetUpDatabase(); err != nil {
		log.Fatal(err)
		panic(err)
	}

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
