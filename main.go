package main

import (
	"fmt"
	"html/template"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	funcMap := template.FuncMap{
		"formatAsDate": formatAsDate,
	}
	//r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	//r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
	r.AddFromFilesFuncs("index", funcMap, "templates/base.html", "templates/index.html")
	r.AddFromFilesFuncs("article", funcMap, "templates/base.html", "templates/index.html", "templates/article.html")
	return r
}

func main() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Html5 Template Engine",
			"now":   time.Now(),
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article", gin.H{
			"title": "Html5 Article Engine",
			"now":   time.Now(),
		})
	})
	router.Run(":8080")
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
