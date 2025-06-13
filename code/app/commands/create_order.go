package commands

import (
	"crud/app/services"
	"crud/domain/model"
	"crud/domain/repository"
)

type OrderService struct {
	orderRepo repository.IOrderRepository
}

func NewOrderHandler(orderRepo repository.IOrderRepository) services.IOrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (h *OrderService) CreateOrderService(order model.Order) error {
	//order := model.NewOrder(cmd.OrderID, cmd.CustomerID, cmd.Total)
	return h.orderRepo.Save(order)
}
