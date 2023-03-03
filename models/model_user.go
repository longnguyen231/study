package models

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type Authen struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type JwtCustomClaims struct {
	Username string `json:"username" :"username"`
	Password bool   `json:"password" :"password"`
	jwt.RegisteredClaims
}
type User struct {
	Id       int    `json:"id" :"id"`
	Name     string `json:"name" :"name"`
	Age      int    `json:"age" :"age"`
	Hometown string `json:"home_town" :"hometown"`
}

func Validate(user User) error {
	if user.Id != 5 {
		return errors.New("Khong phải id")

	}
	return nil
}
