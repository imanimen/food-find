package main

import (
	"github.com/imanimen/foodrate/providers"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(providers.NewConfig),
	)
}