package main

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"os"
)

func parseEnvVariable(name, fallback string) string {
	val := os.Getenv(name)
	if val == "" {
		return fallback
	}
	return val
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.HTMLRender = ginview.New(goview.Config{
		Root:         "views",
		Extension:    ".gohtml",
		Master:       "layouts/master",
		DisableCache: false,
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"placeholder": parseEnvVariable("WEB_PLACEHOLDER", "Web Placeholder"),
			"title":       parseEnvVariable("WEB_TITLE", "Web Placeholder"),
		})
	})

	if err := r.Run(":54321"); err != nil {
		panic(err)
	}
}
