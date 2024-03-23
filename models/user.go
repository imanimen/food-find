package models

type User struct {
	ID  		string  `json:"id"`
	Username 	string  `json:"username"`
	Email 		string  `json:"email"`
	Lat 	  	string  `json:"lat"`
	Long  	  	string  `json:"long"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string 	`json:"updated_at"`
}