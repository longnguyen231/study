package models

import (
	"errors"
)

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
