package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/", func(extc *gin.Context) {
		extc.JSON(http.StatusOK, "ayush here")
	})
	router.Run("127.0.0.1:8080")
}
