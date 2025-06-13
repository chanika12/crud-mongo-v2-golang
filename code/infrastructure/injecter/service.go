package injecter

import (
	"crud/app/services"
	"crud/interfaces/api"
	"github.com/google/wire"
)

func ProvideOrderService(createHandler services.IOrderService) *api.OrderHandler {
	return api.NewOrderHandler(createHandler)
}

var ServiceProviderSet = wire.NewSet(
	ProvideOrderService,
)
