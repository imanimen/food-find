package models

import (

	"math/rand"
)

type OTP struct {
	ID           uint   `gorm:"primaryKey,autoIncrement" json:"id"`
	UserID       string `json:"userId"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Code         string `json:"code"` // Change type to string for OTP code
	CodeExpireAt string `json:"code_expire_at"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (o *OTP) GenerateOTP() string {
	length := 4 
	characters := "123456789"

	randomChars := make([]byte, length)
	for i := 0; i < length; i++ {
		randomChars[i] = characters[rand.Intn(len(characters))]
	}

	code := string(randomChars) // Convert []byte to string for OTP code
	return code
}
