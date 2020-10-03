// handlers.user.go

package main

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func setCookies(c *gin.Context, token, username string, age int) {
	c.SetCookie("token", token, age, "", "", false, true)       // token 30d
	c.SetCookie("username", username, age, "", "", false, true) // token 30d
}

func performLogin(c *gin.Context) {
	username := c.PostForm("usernameLogin")
	password := c.PostForm("passwordLogin")
	check := c.PostForm("checkLogin")

	if username != "" && password != "" && checkFromDB(username, password) {
		token := generateSessionToken()
		if username == "GrandAdmin64" || username == "Admin" {
			sToken := generateSessionToken()
			c.SetCookie("token", token, 300, "", "", false, true)         // token 5m
			c.SetCookie("specialToken", sToken, 300, "", "", false, true) // token 5m
			c.Set("is_logged_in", true)
			c.Set("adminned", true)
		} else {
			if check == "true" {
				setCookies(c, token, username, 2592000) // token 30d
			} else {
				setCookies(c, token, username, 600) // token 10m
			}

			c.Set("is_logged_in", true)
		}

		c.JSON(200, gin.H{
			"message": "Successful login",
		})

	} else {
		c.JSON(200, gin.H{
			"message": "Login Failed: Invalid login or password",
		})
	}
}

// Generates random token
func generateSessionToken() string {
	// hash, _ := HashString("AcCeSs noT deNIeD")
	// return hash
	return strconv.FormatInt(rand.Int63(), 16)
}

// Deletes tokens
func logout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)
	c.SetCookie("specialToken", "", -1, "", "", false, true)
	c.SetCookie("username", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

// Adds new user to DB
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		setCookies(c, token, username, 600) // token 10m
		c.Set("is_logged_in", true)

		showIndexPage(c)
		// render(c, gin.H{
		// 	"title": "Successful registration & Login"}, "login-successful.html")

	} else {
		render(c, gin.H{
			"title":        "Register",
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(),
		}, "register.html")
	}
}

func registerNewUser(username, password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("The password field can't be empty")
	} else if strings.TrimSpace(username) == "" {
		return errors.New("The username field can't be empty")
	} else if !checkUserExist(username) {
		return errors.New("The username isn't available")
	}

	hPass, err := HashString(password)
	if err != nil {
		return err
	}
	u := user{Username: username, Password: hPass}

	addUserToDB(u)

	return nil
}

// Adds new user to DB
func addComment(c *gin.Context) {
	comtitle := c.PostForm("comtitle")
	comment := c.PostForm("comment")

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
