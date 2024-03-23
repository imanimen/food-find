package invokers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/imanimen/foodrate/providers"
	"go.uber.org/fx"
)

func ApiServer(lc fx.Lifecycle, api providers.IApi) *gin.Engine {
	r := gin.Default()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			InitRoutes(r, api)
			go r.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return errors.New("server is down")
		},
	})
	return r
}

func InitRoutes(engine *gin.Engine, api providers.IApi) {
	engine.GET("/", api.Welcome)
}