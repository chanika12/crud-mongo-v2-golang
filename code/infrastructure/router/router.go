package router

import (
	"crud/interfaces/api"
	"github.com/gin-gonic/gin"
)

func StartRouter(orderHandler *api.OrderHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)
	return r
}
