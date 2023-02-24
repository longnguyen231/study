package main

import (
	"log"
	"net/http"
	"projecttest/connect"
	"projecttest/users"

	"github.com/labstack/echo/v4"
)

func main() {
	d := &connect.DB{}
	d.New()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	log.Print(123)
	e.GET("user", users.GetUser)
	e.POST("create", users.CreateUser)
	e.PUT("update/id", users.UpdateUser)
	e.DELETE("/delete", users.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
