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

	// ctx := context.Background()
	db, err := data.SetUpDatabase()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	server := gin.Default()
	templates := template.Must(template.ParseGlob("./templates/*.html"))
	server.SetHTMLTemplate(templates)

	server.Static("/assets", "./assets")
	server.GET("/", func(c *gin.Context) {
		var sayings []data.Saying
		db.Find(&sayings)
		// sayings, err := gorm.G[data.Saying](db)

		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"sayings": sayings,
		})
		log.Println("Wrote the name")
	})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
