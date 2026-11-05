package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"world.com/yappings/data"
	"world.com/yappings/enviornment"
)

func main() {
	shouldMigrate := flag.Bool("migrate", false, "perform automigration syncing the database config with the table setup in the codebase")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	dialector, err := data.GetDialector()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if *shouldMigrate {
		if err := data.Migrate(dialector); err != nil {
			log.Fatal(err)
			panic(err)
		}
		os.Exit(0)
	}

	env := &MainEnv{&enviornment.Environment{Dialector: *dialector}}

	routerEngine := gin.Default()

	multiRender := multitemplate.NewRenderer()
	multiRender.AddFromFiles("main", "templates/index.gohtml", "templates/body.gohtml")
	multiRender.AddFromFiles("error", "templates/index.gohtml", "templates/error-body.gohtml")
	routerEngine.HTMLRender = multiRender

	routerEngine.Static("/assets", "./assets")

	routerEngine.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			errors := c.Errors
			c.HTML(http.StatusInternalServerError, "error", gin.H{"errors": errors})
		}
	})
	routerEngine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(303, "/sayings")
	})
	routerEngine.GET("/sayings", env.sayingsHandler)

	sayingGroup := routerEngine.Group("/saying")
	sayingGroup.GET("/validate", func(ctx *gin.Context) {
	})
	sayingGroup.POST("/", func(c *gin.Context) {
	})
	sayingGroup.PUT("/", func(c *gin.Context) {
	})

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routerEngine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

type MainEnv struct {
	*enviornment.Environment
}

func (env *MainEnv) sayingsHandler(ctx *gin.Context) {
	db, err := gorm.Open(env.Dialector, &gorm.Config{})

	if err != nil {
		ctx.Error(err)
		return
	}

	sayings, err := gorm.G[data.Saying](db).Find(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "main", gin.H{
		"sayings": sayings,
	})
}
