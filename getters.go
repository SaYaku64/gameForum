package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getCurrentTime() string {
	now := time.Now()
	return now.Format("02 Jan 06 15:04")
}

// Gets and shows specific user article
func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Gets specific User article by ID from DB
func getArticleByID(id int) (*article, error) {
	for _, a := range getArticleFromDB() {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
