package di

import (
	"github.com/google/wire"

	"github.com/firstics/pos-be/pkg/api/handler"
	"github.com/firstics/pos-be/pkg/usecase"
)

var ExampleSet = wire.NewSet(
	usecase.NewExampleUsecase,
	handler.NewExampleHandler,
)
