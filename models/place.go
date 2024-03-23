package models


type Place struct {
	ID        uint    `gorm:"primaryKey,autoIncrement" json:"id"`
	Lat 	  string  `json:"lat"`
	Long  	  string   `json:"long"`
	CreatedAt string  `json:"created_at"`
}
