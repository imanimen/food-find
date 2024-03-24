package providers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/imanimen/foodrate/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDatabase interface {
	/*
	* Auth Methods
	*/
	sendOTP(email string) (string, string, error)
	verifyCode(email, code string) (string, error)
}

type Database struct {
	Connection *gorm.DB
	Config     IConfig
}

func NewDatabase(config IConfig) IDatabase {
	dsn := config.Get("dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting db")
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.OTP{})
	db.AutoMigrate(&models.Place{})

	return &Database{
		Connection: db,
		Config:     config,
	}
}


func (database *Database) sendOTP(email string) (string, string, error) {
	var user models.User
	err := database.Connection.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newUser := models.User{
				ID:    uuid.New().String(),
				Email: email,
				Lat: "0",
				Long: "0",
				CreatedAt: time.Now().Format(time.RFC3339),
				UpdatedAt: time.Now().Format(time.RFC3339),
			}

			if err := database.Connection.Create(&newUser).Error; err != nil {
				return "", "", err
			}

			user = newUser
		} else {
			return "", "", err
		}
	}

	var otp models.OTP
	err = database.Connection.Where("user_id = ?", user.ID).First(&otp).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return "", "", err
		}
	}

	if otp.ID == 0 {
		newCode := otp.GenerateOTP()
		expirationTime := time.Now().Add(time.Minute * 5).Format(time.RFC3339)

		otp = models.OTP{
			UserID:       user.ID,
			Email:        user.Email,
			PhoneNumber:  "", // Add phone number if needed
			Code:         newCode,
			CodeExpireAt: expirationTime,
			CreatedAt:    time.Now().Format(time.RFC3339),
			UpdatedAt:    time.Now().Format(time.RFC3339),
		}

		if err := database.Connection.Create(&otp).Error; err != nil {
			return "", "", err
		}

		return newCode, expirationTime, nil
	}

	newCode := otp.GenerateOTP()
	expirationTime := time.Now().Add(time.Minute * 5).Format(time.RFC3339)

	if err := database.Connection.Model(&otp).Updates(models.OTP{Code: newCode, CodeExpireAt: expirationTime}).Error; err != nil {
		return "", "", err
	}

	return newCode, expirationTime, nil
}

func (database *Database) verifyCode(email, code string) (string, error) {
	var otp models.OTP
	err := database.Connection.Where("email = ? AND code = ?", email, code).First(&otp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("OTP record not found")
		}
		return "", err
	}

	// Update the OTP record with null values for Code and CodeExpireAt
	if err := database.Connection.Model(&otp).UpdateColumn("code", nil).Error; err != nil {
		return "", err
	}

	// // Generate JWT token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.CustomClaims{Email: email})
	// signedToken, err := token.SignedString([]byte("lajksdalksjdalksjd"))
	// if err != nil {
	// 	return "", err
	// }

	return otp.UserID, nil
}
