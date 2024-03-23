package models


type Place struct {
	ID        		uint    `gorm:"primaryKey,autoIncrement" json:"id"`
	Title 	  		string  `json:"title"`
	Description 	string  `json:"description"`
	Logo 			string  `json:"logo"`
	Lat 	  		string  `json:"lat"`
	Long  	  		string   `json:"long"`
	CreatedAt 		string  `json:"created_at"`
}
