package api

import (
	"crud/app/commands"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	createHandler *commands.CreateOrderHandler
}

func NewOrderHandler(h *commands.CreateOrderHandler) *OrderHandler {
	return &OrderHandler{createHandler: h}
}

type CreateOrderRequest struct {
	OrderID    string  `json:"order_id" binding:"required"`
	CustomerID string  `json:"customer_id" binding:"required"`
	Total      float64 `json:"total" binding:"required"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := commands.CreateOrderCommand{
		OrderID:    req.OrderID,
		CustomerID: req.CustomerID,
		Total:      req.Total,
	}

	if err := h.createHandler.Handle(c.Request.Context(), cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
