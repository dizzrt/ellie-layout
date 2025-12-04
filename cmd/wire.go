//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dizzrt/ellie"
	"github.com/dizzrt/ellie-layout/internal/application"
	"github.com/dizzrt/ellie-layout/internal/conf"
	"github.com/dizzrt/ellie-layout/internal/domain"
	"github.com/dizzrt/ellie-layout/internal/handler"
	"github.com/dizzrt/ellie-layout/internal/infra"
	"github.com/dizzrt/ellie-layout/internal/server"
	"github.com/google/wire"
)

func wireApp() (*ellie.App, func(), error) {
	panic(wire.Build(
		conf.ProviderSet,
		infra.ProviderSet,
		domain.ProviderSet,
		application.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
		newApp,
	))
}
