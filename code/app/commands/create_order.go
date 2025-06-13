package commands

import (
	"context"
	"crud/domain/model"
	"crud/domain/repository"
)

type CreateOrderCommand struct {
	OrderID    string
	CustomerID string
	Total      float64
}

type CreateOrderHandler struct {
	OrderRepo repository.OrderRepository
}

func NewCreateOrderHandler(orderRepo repository.OrderRepository) *CreateOrderHandler {
	return &CreateOrderHandler{OrderRepo: orderRepo}
}

func (h *CreateOrderHandler) Handle(ctx context.Context, cmd CreateOrderCommand) error {
	order := model.NewOrder(cmd.OrderID, cmd.CustomerID, cmd.Total)
	return h.OrderRepo.Save(ctx, order)
}
