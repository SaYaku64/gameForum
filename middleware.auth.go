// middleware.auth.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

////////////////////////////////////////////
// Checking user status
////////////////////////////////////////////
func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
			render(c, gin.H{
				"title": "Home Page",
			}, "index.html")
		}
	}
}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
			render(c, gin.H{
				"title": "Home Page",
			}, "index.html")
		}
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" { // && CheckPasswordHash("AcCeSs noT deNIeD", token)
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
