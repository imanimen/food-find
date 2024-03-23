package providers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type IApi interface {
	Welcome(c *gin.Context)
}

type Api struct {
	Config      IConfig
	Database    IDatabase
}

func NewApi(config IConfig, database IDatabase) IApi {
	return &Api{
		Config:      config,
		Database:    database,
	}
}


func (api *Api) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": api.Config.Get("apiVersion"),
		"data":    "Hello World",
	})

}


