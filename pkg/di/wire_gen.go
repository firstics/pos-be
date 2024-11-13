// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/firstics/pos-be/pkg/api"
	"github.com/firstics/pos-be/pkg/api/handler"
	"github.com/firstics/pos-be/pkg/api/middleware"
	"github.com/firstics/pos-be/pkg/config"
	"github.com/firstics/pos-be/pkg/driver"
	"github.com/firstics/pos-be/pkg/usecase"
)

// Injectors from wire.go:

func InitializeApp(config2 config.Config) (*api.ServerHTTP, error) {
	logger := driver.NewLogger(config2)
	errorHandler := middleware.NewErrorHandler(logger)
	middlewares := &api.Middlewares{
		ErrorHandler: errorHandler,
	}
	exampleUsecase := usecase.NewExampleUsecase(logger)
	exampleHandler := handler.NewExampleHandler(exampleUsecase)
	handlers := api.Handlers{
		ExampleHandler: exampleHandler,
	}
	serverHTTP := api.NewServerHTTP(middlewares, handlers, logger, config2)
	return serverHTTP, nil
}