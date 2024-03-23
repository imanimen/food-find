package main

import (
	"github.com/imanimen/foodrate/invokers"
	"github.com/imanimen/foodrate/providers"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(providers.NewConfig, providers.NewDatabase, providers.NewApi),
		fx.Invoke(invokers.ApiServer),
	).Run()
}
