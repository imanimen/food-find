package providers

import (
	"net/mail"
	"regexp"
)



type IValidations interface {
	IsValidEmail(string) bool
	IsValidMobile(string) bool
}

type Validations struct {
	Email 			string `form:"email" binding:"required,email"`
	PhoneNumber  	string `form:"mobile" binding:"required"`
}

// IsValidEmail validates that the given email address is a valid format.
// It checks that the email is not empty, has a minimum length of 3 characters,
// and passes go's mail.ParseAddress validation.
func (v Validations) IsValidEmail(email string) bool {

	if email == "" {
		return false
	}

	if len(email) < 3 {
		return false
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return true
}

// IsValidMobile validates if the given phoneNumber is a valid
// Iranian mobile phone number format. It uses a regex to match
// the common Iranian mobile number format. Returns true if valid,
// false otherwise.
func (v Validations) IsValidMobile(phoneNumber string) bool {
	//  for Iranian phone numbers
	iranRegex := `^(\+98|0)?9\d{9}$`
	re := regexp.MustCompile(iranRegex)

	return re.MatchString(phoneNumber)
}

func NewValidations() IValidations {
	return &Validations{}
}