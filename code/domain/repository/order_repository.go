package repository

import (
	"context"
	"crud/domain/model"
)

type OrderRepository interface {
	Save(ctx context.Context, order *model.Order) error
}
