package providers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IApi interface {
	Welcome(c *gin.Context)

	/*
	* Authorization API
	*/
	SendCode(c *gin.Context)
	VerifyCode(c *gin.Context)

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


// send otp

func (api *Api) SendCode(c *gin.Context) {
	email := c.PostForm("email")
	code, expireAt, err := api.Database.sendOTP(email)
	// TODO: error handle the channel
	// go func() {
	// 	services.Call(api.Config.Get("NOTIFY_API_URL") + "/v1/notify", "POST", payloadData)
	// }()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"version": api.Config.Get("apiVersion"),
		"data": gin.H{
			"code":      code,
			"expire_at": expireAt,
		},
	})

}


func (api *Api) VerifyCode(c *gin.Context) {
	email := c.PostForm("email")
	code  := c.PostForm("code")
	result, err := api.Database.verifyCode(email, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"version": api.Config.Get("apiVersion"),
		"data":    result,
	})


}