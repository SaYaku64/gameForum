// handlers.user.go

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quickemailverification/quickemailverification-go"
)

// func plusThirty() time.Time {
// 	return time.Now().Add(time.Minute * 30)
// }

func performLogin(c *gin.Context) {
	username := c.PostForm("usernameLogin")
	password := c.PostForm("passwordLogin")
	check := c.PostForm("checkLogin")

	if username != "" && password != "" {
		ok, curUser := checkFromDB(username, password)

		if ok {
			c.Set("is_logged_in", true)
			token := generateSessionToken()

			if curUser.AdminStatus {

				sToken := curUser.AdminToken
				c.SetCookie("token", token, 300, "", "", false, true)         // token 5m
				c.SetCookie("specialToken", sToken, 300, "", "", false, true) // token 5m
				c.Set("adminned", true)
				addToken(curUser.Username, token, false)
			} else {
				if check == "true" {
					c.SetCookie("token", token, 12592000, "", "", false, true)
					addToken(curUser.Username, token, true)
				} else {
					c.SetCookie("token", token, 600, "", "", false, true) // token 10m
					addToken(curUser.Username, token, false)
				}
			}

			curUser.addUserToCache()

			c.JSON(200, gin.H{
				"message": "Successful login",
			})
		}
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
	return fmt.Sprint(time.Now().UnixNano())
}

// Deletes tokens
func logout(c *gin.Context) {

	u, err := getCurrentUser(c)
	if err != nil {
		log.Println(err)
	} else {
		u.delUserFromCache(u.Token.Name)
	}
	delToken(u.Username)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.SetCookie("specialToken", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (u user) addUserToCache() {
	ActiveUsers.RLock()
	ActiveUsers.m[u.Token.Name] = u
	ActiveUsers.RUnlock()
}

func (u user) delUserFromCache(name string) {
	ActiveUsers.RLock()
	delete(ActiveUsers.m, name)
	ActiveUsers.RUnlock()
}

// Adds new user to DB
func register(c *gin.Context) {
	email := c.PostForm("email")

	result := checkEmailValidation(email)
	if result == "valid" {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if token, err := registerNewUser(email, username, password); err == nil {
			c.SetCookie("token", token, 600, "", "", false, true) // token 10m
			c.Set("is_logged_in", true)

			// showIndexPage(c)
			render(c, gin.H{
				"title": "Home page"}, "index.html")
		} else {
			render(c, gin.H{
				"title":        "Register",
				"ErrorTitle":   "Registration Failed",
				"ErrorMessage": err.Error(),
			}, "register.html")
		}
	} else {
		render(c, gin.H{
			"title":        "Register",
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": "Invalid email adress",
		}, "register.html")
	}

}

func checkEmailValidation(email string) string {
	qev := quickemailverification.CreateClient("929f903df3aeeae441e711bcbd1d743d42cd1f0364a8d78a3fed131af710")
	// Need to use Verify instead Sandbox in production
	response, err := qev.Sandbox(email) // Email address which need to be verified
	if err != nil {
		log.Println(err)
		return "Validation failed"
	}
	return response.Result
}

func registerNewUser(email, username, password string) (string, error) {
	if strings.TrimSpace(password) == "" {
		return "", errors.New("The password field can't be empty")
	} else if strings.TrimSpace(email) == "" {
		return "", errors.New("The email adress field can't be empty")
	} else if !checkEmailExist(email) {
		return "", errors.New("The email is already used")
	} else if strings.TrimSpace(username) == "" {
		return "", errors.New("The username field can't be empty")
	} else if !checkUserExist(username) {
		return "", errors.New("The username is already used")
	}

	hPass, err := HashString(password)
	if err != nil {
		return "", err
	}
	u := user{Email: email, Username: username, Password: hPass, Token: token{Name: generateSessionToken()}}

	err = addUserToDB(u)
	if err != nil {
		return "", err
	}

	u.addUserToCache()

	return u.Token.Name, nil
}
