package server

import (
	"github.com/sellist/personal-website-api/internal/testing"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func Run() {
	address := os.Getenv("PORT")
	remoteAddr := os.Getenv("REMOTE_ADDR")
	dyno := os.Getenv("DYNO")

	if dyno == "" {
		log.Println("No dyno detected, running locally")
		address = "localhost:8080"
	} else {
		address = ":" + address
	}

	if address == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	if len(remoteAddr) > 0 {
		err := router.SetTrustedProxies([]string{remoteAddr})
		if err != nil {
			log.Fatal(err)
		}
	}

	router.GET("/", testing.HandleTest)

	runError := router.Run(address)
	if runError != nil {
		log.Fatal(runError)
	}
}
