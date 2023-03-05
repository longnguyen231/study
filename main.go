package main

import (
	"net/http"
	"projecttest/connect"
	"projecttest/middleware"
	"projecttest/users"

	"github.com/labstack/echo/v4"
)

func main() {
	d := &connect.DBSql{}
	d.New()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("user", users.GetUser)
	e.POST("create", users.CreateUser)
	e.PUT("update/id", users.UpdateUser)
	e.DELETE("/delete", users.DeleteUser)
	e.GET("/login", users.Login)
	e.GET("/protected", users.Protected, middleware.VerifyJWT)
	e.Logger.Fatal(e.Start(":1323"))
}
