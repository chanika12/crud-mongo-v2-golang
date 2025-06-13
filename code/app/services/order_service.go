package services

import (
	"crud/domain/model"
)

type IOrderService interface {
	CreateOrderService(order model.Order) error
}
