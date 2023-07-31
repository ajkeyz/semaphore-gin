package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")

}

func getArticle(c *gin.Context) {
	// check article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// check that the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// call HTML method of Context to render a template
			c.HTML(
				// set the HTTP status to 200 OK
				http.StatusOK,
				// use the article.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			// if article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// if we find an invalid article ID in URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
