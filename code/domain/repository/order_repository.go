package repository

import (
	"crud/domain/model"
)

type IOrderRepository interface {
	Save(order model.Order) error
}
