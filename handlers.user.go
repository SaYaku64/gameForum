// handlers.user.go

package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func performLogin(c *gin.Context) {
	username := c.PostForm("usernameLogin")
	password := c.PostForm("passwordLogin")

	if checkFromDB(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		// c.JSON(200, gin.H{
		// 	"message": "Successful login",
		// })
		//showIndexPage(c)
		// render(c, gin.H{
		// 	"title": "Successful Login"}, "login-successful.html")

	} else {
		c.JSON(200, gin.H{
			"message": "Login Failed: Invalid login or password",
		})
		// render(c, gin.H{
		// 	"title":        "Home Page",
		// 	"ErrorTitle":   "Login Failed",
		// 	"ErrorMessage": "Invalid login or password",
		// }, "index.html")

		// articles := getArticleFromDB()
		// render(c, gin.H{
		// 	"title":        "Home Page",
		// 	"payload":      articles,
		// 	"ErrorTitle":   "Login Failed",
		// 	"ErrorMessage": "Invalid login or password",
		// }, "index.html")

		// c.HTML(http.StatusBadRequest, "index.html", gin.H{
		// 	"ErrorTitle":   "Login Failed",
		// 	"ErrorMessage": "Invalid login or password"})
	}
}

func generateSessionToken() string {
	// hash, _ := HashString("AcCeSs noT deNIeD")
	// return hash
	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
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
