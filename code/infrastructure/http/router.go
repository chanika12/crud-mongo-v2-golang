package http

import (
	"crud/interfaces/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter(orderHandler *api.OrderHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)
	return r
}
