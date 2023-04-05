package models

import "github.com/golang-jwt/jwt"

type BaseClient struct {
	Base
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	Email     string `form:"email" json:"email" gorm:"unique"`
	Address   string `form:"address" json:"address"`
	Password  string `form:"password" json:"password"`
}
type Clients struct {
	BaseClient
	ProfilePhoto string `form:"profile_photo" json:"profile_photo" bson:",omitempty" `
}

func (m Clients) TableName() string {
	return "clients"
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClientRequestResponse struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ImageUrl     string `json:"image_url"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequestResponse struct {
	AcceessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignedDetails struct {
	Email string
	jwt.StandardClaims
}

type ClientResponse struct {
	Base
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	Email     string `form:"email" json:"email" gorm:"unique"`
	Address   string `form:"address" json:"address"`
	ProfilePhoto string `form:"profile_photo" json:"profile_photo" bson:",omitempty" `
}
