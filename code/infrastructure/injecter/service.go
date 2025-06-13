package injecter

import (
	"crud/app/commands"
	"crud/app/services"
	"crud/domain/repository"
	"github.com/google/wire"
)

func ProvideOrderService(createHandler repository.IOrderRepository) services.IOrderService {
	return commands.NewOrderService(createHandler)
}

var ServiceProviderSet = wire.NewSet(
	ProvideOrderService,
)
