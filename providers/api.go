package providers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IApi defines the interface for the API provider.
// It contains endpoints for user authorization and profile management.
type IApi interface {
	Welcome(c *gin.Context)

	/*
	* Authorization API
	*/
	SendCode(c *gin.Context)
	VerifyCode(c *gin.Context)

	/*
	* Profile Area
	*/
	Me(c *gin.Context)
	UpdateProfile(c *gin.Context)
}

// Api holds the application dependencies.
type Api struct {
	Config   IConfig
	Database IDatabase
	Validations IValidations
}

// NewApi creates a new Api instance with the given config and database.
func NewApi(config IConfig, database IDatabase, validation IValidations) IApi {
	return &Api{
		Config:   config,
		Database: database,
		Validations: validation,
	}
}

// Welcome handles the /welcome API endpoint. It returns a simple JSON
// response with the API version and a greeting message.
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
	if !api.Validations.IsValidEmail(email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid email format",
		})
		return
	}
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
// and returns the result and any error
func (api *Api) VerifyCode(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	if !api.Validations.IsValidEmail(email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid email format",
		})
		return
	}
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
// UpdateProfile updates the profile information for the user with the given ID.
// It takes the ID, latitude, longitude, and username from the request and updates
// the corresponding fields in the user object. It returns the updated user object
// on success. Returns 422 if ID is missing. Returns 500 on error retrieving user
// or updating profile.
func (api *Api) UpdateProfile(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "id is required",
		})
		return
	}

	// Parse update data from request body
	lat := c.PostForm("lat")
	long := c.PostForm("long")
	username := c.PostForm("username")

	// Validate input
	if lat == "" || long == "" || username == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "lat, long, and username are required fields",
		})
		return
	}

	user, err := api.Database.getUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user data",
		})
		return
	}

	user.Lat = lat
	user.Long = long
	user.Username = username

	user, err = api.Database.updateProfile(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"version": api.Config.Get("apiVersion"),
		"data":    user,
	})
}
