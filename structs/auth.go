package structs

import (
	"99live-cms-golang-api/models"

	jwtLib "github.com/dgrijalva/jwt-go"
)

type LoginModel struct {
	Token TokenModel  `json:"token"`
	User  models.User `json:"user"`
}

type TokenModel struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type JwtClaims struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
	jwtLib.StandardClaims
}
