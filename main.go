package main

import (
	"context"
	"database/sql"
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

	routerEngine.HTMLRender = createRenderer()
	routerEngine.Static("/assets", "./assets")

	routerEngine.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			errors := c.Errors
			c.HTML(http.StatusInternalServerError, "error", gin.H{"errors": errors})
		}
	})

	routerEngine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/sayings")
	})
	routerEngine.GET("/sayings", env.sayingsHandler)

	sayingGroup := routerEngine.Group("/saying")
	sayingGroup.GET("/validate", func(ctx *gin.Context) {
	})
	sayingGroup.POST("/", env.handleSayingPost)
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

func createRenderer() multitemplate.Renderer {
	multiRender := multitemplate.NewRenderer()
	multiRender.AddFromFiles("main", "templates/index.gohtml", "templates/body.gohtml", "templates/sayingsList.gohtml")
	multiRender.AddFromFiles("error", "templates/index.gohtml", "templates/error-body.gohtml")
	multiRender.AddFromFiles("sayingsList", "templates/sayingsList.gohtml")
	return multiRender
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

func (env *MainEnv) handleSayingPost(ctx *gin.Context) {

	original := ctx.PostForm("original")
	translated := ctx.PostForm("translated")
	meaning := ctx.PostForm("meaning")
	whentouse := ctx.PostForm("whentouse")

	whentouseNull := sql.NullString{String: whentouse}

	saying := data.Saying{
		Original:    original,
		Translation: translated,
		Meaning:     meaning,
		WhenToUse:   whentouseNull,
	}

	db, err := gorm.Open(env.Dialector, &gorm.Config{})

	if err != nil {
		ctx.Error(err)
		return
	}

	bgCtx := context.Background()
	err = gorm.G[data.Saying](db).Create(bgCtx, &saying)

	if err != nil {
		ctx.Error(err)
		return
	}

	newSayings, err := gorm.G[data.Saying](db).Find(bgCtx)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "sayingsList", gin.H{
		"sayings": newSayings,
	})
}
