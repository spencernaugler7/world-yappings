package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
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

	templates := template.Must(template.ParseGlob("./templates/*.html"))

	router := gin.Default()
	router.SetHTMLTemplate(templates)
	router.Static("/assets", "./assets")

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(303, "/sayings")
	})
	router.GET("/sayings", func(c *gin.Context) {

		sayings, err := gorm.G[data.Saying](db).Find(c)

		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"sayings": sayings,
		})
		log.Println("Wrote the name")
	})

	saying := router.Group("/saying")
	saying.GET("/validate", func(ctx *gin.Context) {

	})
	saying.POST("/", func(c *gin.Context) {

	})
	saying.PUT("/", func(c *gin.Context) {

	})

	server := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
