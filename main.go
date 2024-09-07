package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {

	router := gin.Default()

	router.GET("/", func(extc *gin.Context) {
		extc.JSON(http.StatusOK, "ayush here")
	})

	// Azure App Service sets the port as an Environment Variable
	// This can be random, so needs to be loaded at startup
	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	router.Run("127.0.0.1:" + portt)
}
