package main

import (
	"crud/infrastructure/config"
	"crud/infrastructure/injecter"
	"crud/infrastructure/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	config.LoadConfig()
	initAdapter := injecter.InitializeAdapter()

	r.GET("api/bill-manager/v1/bill-manager/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.StartRouter(initAdapter)

	r.Run(":3000")
}
