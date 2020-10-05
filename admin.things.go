package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Shows admin-panel
func showAdminPanelPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Admin Panel"}, "admin-panel.html")
}

// Deletes things, that were chosen in admin-panel
func delThisShit(c *gin.Context) {
	delElem := c.PostForm("delElem")
	delName := c.PostForm("delName")

	if strings.TrimSpace(delElem) == "" || strings.TrimSpace(delName) == "" {
		c.JSON(200, gin.H{
			"message": "Error: Empty field!",
		})
	} else {
		delElem = strings.TrimSpace(delElem)
		var collection mongo.Collection
		var filter bson.M
		if strings.ToLower(delElem) == "users" {
			collection = *Client.Database("courses").Collection("users")
			filter = bson.M{"username": delName}
		} else if strings.ToLower(delElem) == "articles" {
			collection = *Client.Database("courses").Collection("articles")
			filter = bson.M{"title": delName}
		}

		// update := bson.M{
		// 	"$set": bson.M{"comment": []comment{}},
		// }

		// deleteResult, err := collection.UpdateOne(context.TODO(), filter, update)
		deleteResult, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"message": "Error: Failed to delete!",
			})
		}
		log.Println(deleteResult)
		c.JSON(200, gin.H{
			"message": "Successfully deleted",
		})
	}

}

////////////////////////////////////////////
// Checking admin status
////////////////////////////////////////////
func ensureAdminned() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminnedInterface, _ := c.Get("adminned")
		adminnedIn := adminnedInterface.(bool)

		if !adminnedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
			render(c, gin.H{
				"title": "Home Page",
			}, "index.html")
		}
	}
}

func setAdminStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("specialToken"); err == nil || token != "" {
			c.Set("adminned", true)
		} else {
			c.Set("adminned", false)
		}
	}
}
