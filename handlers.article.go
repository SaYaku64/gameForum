// handlers.article.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(c, title, content); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// Creating article
func createNewArticle(c *gin.Context, title, content string) (*article, error) {
	time := getCurrentTime()
	name := ""
	if token, err := c.Cookie("username"); err == nil || token != "" {
		name = token
	}

	a := article{ID: len(getArticleFromDB()) + 1, Title: title, Content: content, Time: time, Name: name, Comment: []comment{}}

	insertArticleToDB(a)

	return &a, nil

}
