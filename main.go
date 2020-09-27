// main.go

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var router *gin.Engine

// Client for MongoDB
var Client = connDB()

func main() {

	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")

	initializeRoutes()

	err := router.Run()
	if err != nil {
		log.Println(err)
	}

	disconnDB(Client)
}

func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, ok := c.Get("is_logged_in")
	if ok != true {
		log.Println("Error in getting 'is_logged_in' parameter")
		return
	}
	data["is_logged_in"], ok = loggedInInterface.(bool)
	if ok != true {
		log.Println("LoggedInInterface isn't bool")
		return
	}

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func connDB() mongo.Client {
	// Создаём клиента БД
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Println(err)
	}

	// Создаём соединение
	err = client.Connect(context.TODO())
	if err != nil {
		log.Println(err)
	}

	// Проверяем соединение
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}

	return *client
}

func disconnDB(client mongo.Client) {
	// Закрываем соединение с БД
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Println(err)
	}
}
