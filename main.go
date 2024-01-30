package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func Test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello world!")
}

func main() {
	port := os.Getenv("PORT")
	addr := os.Getenv("REMOTE_ADDR")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{addr})
	router.GET("/", Test)

	router.Run(":" + port)
}
