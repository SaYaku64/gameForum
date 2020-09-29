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

//////////////////////////////////////////////////////////////////////////////
// TUTORIAL
//////////////////////////////////////////////////////////////////////////////
// Gets and shows specific Tutorial article
func getTutorialArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getTutorialByID(articleID); err == nil {
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

// Gets specific Tutorial by ID from cache of Tutorials
func getTutorialByID(id int) (*article, error) {
	for _, a := range Tutorials {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

//////////////////////////////////////////////////////////////////////////////
// QUESTION
//////////////////////////////////////////////////////////////////////////////
// Gets and shows specific Question article
func getFAQArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getFAQByID(articleID); err == nil {
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

// Gets specific Question by ID from cache of Questions
func getFAQByID(id int) (*article, error) {
	for _, a := range FAQs {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

//////////////////////////////////////////////////////////////////////////////
// ARTICLE
//////////////////////////////////////////////////////////////////////////////
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
