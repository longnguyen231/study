package models

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Username string `json:"username" :"username"`
	Password bool   `json:"password" :"password"`
	jwt.RegisteredClaims
}
type User struct {
	Id       int    `json:"id" :"id" :"id"`
	Name     string `json:"name" :"name" :"name"`
	Age      int    `json:"age" :"age" :"age"`
	Hometown string `json:"home_town" :"hometown" :"hometown"`
	Password string `json:"password" :"password"`
	Username string `json:"username" :"username"`
}
