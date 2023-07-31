// main.go

package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// set router as default one provided by Gin
	router = gin.Default()
	//load templates at start to avoid loading from disk each time
	router.LoadHTMLGlob("templates/*")

	//define route handlers
	router.GET("/", showIndexPage)
	 // Handle GET requests at /article/view/some_article_id
	 router.GET("/article/view/:article_id", getArticle)

	// Start serving the application
	router.Run(":8080")

}
