package models

import "os"

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Lat       string `json:"lat"`
	Long      string `json:"long"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Avatar    string `json:"avatar"`
}

func (u *User) GetAvatar() string {
	return os.Getenv("CDN_URL")+u.Avatar
}