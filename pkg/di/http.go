package di

import (
	"github.com/google/wire"

	"github.com/pos-be/pkg/api"
	"github.com/pos-be/pkg/api/middleware"
)

var HttpSet = wire.NewSet(
	api.NewServerHTTP,
	middleware.NewErrorHandler,
	wire.Struct(new(api.Middlewares), "*"),
	wire.Struct(new(api.Handlers), "*"),
)
