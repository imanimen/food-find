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
	Me(c *gin.Context)

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


// SendCode sends an OTP code to the provided email address and returns
// the code, expiration time, and any error.

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

// VerifyCode verifies the provided OTP code against the email address
// and returns the result and any error.

func (api *Api) VerifyCode(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
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


// Me retrieves the user with the given ID from the database
// and returns it in the response. Returns status code 200 and the
// user data on success. Returns status code 422 if the ID is missing.
// Returns status code 500 on error.
func (api *Api) Me(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "id is required",
		})
	}
	user, _ := api.Database.getUserByID(id)
	c.JSON(http.StatusOK, gin.H{
		"version": api.Config.Get("apiVersion"),
		"data":    user,
	})
}
