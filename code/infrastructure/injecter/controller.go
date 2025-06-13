package injecter

import (
	"crud/app/services"
	"crud/interfaces/api"
	"github.com/google/wire"
)

func ProvideOrderController(createHandler services.IOrderService) *api.OrderHandler {
	return api.NewOrderHandler(createHandler)
}

var ControllerProviderSet = wire.NewSet(
	ProvideOrderController,
)
