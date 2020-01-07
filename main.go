package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	initializeRoutes()
	router.Run()	// Runs on default port 8080
}

// This function will initialize all the routes
func initializeRoutes() {
	router.GET("/hello", helloHandler)
	router.GET("/greet/:name", greetHandler)
}

// Returns a HelloWorld json message
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "message": "Hello World!"})
}

// Returns a greeting json message
func greetHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{ "greet-message": "Hello " + name + "!" })
}
