package main

import (
	"github.com/imanimen/foodrate/invokers"
	"github.com/imanimen/foodrate/providers"
	"go.uber.org/fx"
)

// main is the entry point for the application. It initializes the dependency injection
// container using fx.New, registers providers and invokers, and starts the server.
func main() {
	fx.New(
		fx.Provide(providers.NewConfig, providers.NewDatabase, providers.NewApi),
		fx.Invoke(invokers.ApiServer),
	).Run()
}
