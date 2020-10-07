// middleware.auth.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// // Function to substract one time from anouther
// func subtractTime(time1, time2 time.Time) float64 {
// 	diff := time1.Sub(time2).Seconds()
// 	return diff
// }

////////////////////////////////////////////
// Checking user status
////////////////////////////////////////////
func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, ok := c.Get("is_logged_in")
		if ok {
			loggedIn := loggedInInterface.(bool)
			if !loggedIn {
				c.AbortWithStatus(http.StatusUnauthorized)
				render(c, gin.H{
					"title": "Home Page",
				}, "index.html")
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			render(c, gin.H{
				"title": "Home Page",
			}, "index.html")
		}

	}
}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, ok := c.Get("is_logged_in")
		if ok {
			loggedIn := loggedInInterface.(bool)
			if loggedIn {
				c.AbortWithStatus(http.StatusUnauthorized)
				render(c, gin.H{
					"title": "Home Page",
				}, "index.html")
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			render(c, gin.H{
				"title": "Home Page",
			}, "index.html")
		}

	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {

		// if token, err := c.Cookie("token"); err == nil || token != "" {
		// 	c.Set("is_logged_in", true)
		// } else {
		// 	c.Set("is_logged_in", false)
		// }

		// log.Println(ActiveUsers)
		if token, err := c.Cookie("token"); err == nil || token != "" { // take cookie
			//ActiveUsers.RLock()
			if u, ok := ActiveUsers.m[token]; ok { // take user from cache
				//ActiveUsers.RUnlock()
				if u.Token.Endless {
					c.Set("is_logged_in", true)
				} else {
					// diff := subtractTime(u.Token.EndTime, time.Now()) // checking if user
					// if diff > 0 {
					c.Set("is_logged_in", true)
					c.SetCookie("token", token, 600, "", "", false, true) // token 10m
					// }
				}
			}
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
