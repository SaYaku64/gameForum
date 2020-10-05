// handlers.article.go

package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if _, err := createNewArticle(c, title, content); err == nil {
		c.JSON(200, gin.H{
			"message": "Successful creation",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Creation failed",
		})
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

// Removes all comments on post
func deleteArticle(c *gin.Context) {
	title := c.PostForm("title")
	articleAuthor := c.PostForm("authorName")

	username := ""
	if token, err := c.Cookie("username"); err == nil || token != "" {
		username = token
	}

	if username == articleAuthor || username == "" {
		if err := deleteArticleFromDB(title); err == nil {
			c.JSON(200, gin.H{
				"message": "Successfully deleted",
			})
		} else {
			c.JSON(200, gin.H{
				"message": err,
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "You are not author of the article!",
		})
	}
}

// Adds comment on post
func addComment(c *gin.Context) {
	comtitle := c.PostForm("comtitle")
	comment := c.PostForm("comment")
	if strings.TrimSpace(comtitle) == "" || strings.TrimSpace(comment) == "" {
		c.JSON(200, gin.H{
			"message": "Error: Empty field!",
		})
	} else {
		time := getCurrentTime()
		name := ""

		if token, err := c.Cookie("username"); err == nil || token != "" {
			name = token
		}

		if err := commentToDB(comtitle, comment, time, name); err == nil {
			c.JSON(200, gin.H{
				"message": "Successfully added",
			})

		} else {
			c.JSON(200, gin.H{
				"message": "Failed adding comment",
			})
		}
	}
}

// Removes all comments on post
func removeComment(c *gin.Context) {
	title := c.PostForm("title")
	articleAuthor := c.PostForm("authorName")

	username := ""
	if token, err := c.Cookie("username"); err == nil || token != "" {
		username = token
	}

	if username == articleAuthor || username == "" {
		if err := delComment(title); err == nil {
			c.JSON(200, gin.H{
				"message": "Successfully removed",
			})

		} else {
			c.JSON(200, gin.H{
				"message": err,
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "You are not author of the article!",
		})
	}
}
