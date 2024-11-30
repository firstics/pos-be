//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/pos-be/pkg/api"
	"github.com/pos-be/pkg/config"
)

func InitializeApp(config config.Config) (*api.ServerHTTP, error) {
	wire.Build(ExampleSet, LogSet, HttpSet)

	return &api.ServerHTTP{}, nil
}
