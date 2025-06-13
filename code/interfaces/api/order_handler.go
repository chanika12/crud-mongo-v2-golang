package api

import (
	"crud/app/services"
	"crud/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	createHandler services.IOrderService
}

func NewOrderHandler(h services.IOrderService) *OrderHandler {
	return &OrderHandler{createHandler: h}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req model.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.createHandler.CreateOrderService(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
