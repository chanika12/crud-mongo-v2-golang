package injecter

import (
	"crud/interfaces/api"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	RepositoryProviderSet,
	ServiceProviderSet,
	ControllerProviderSet,
)

func InitializeAdapter() *api.OrderHandler {
	wire.Build(ProviderSet)
	return &api.OrderHandler{}
}
