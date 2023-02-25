package jwt

import "github.com/golang-jwt/jwt/v4"

type MyClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte("Douyin_Sixwithone")
