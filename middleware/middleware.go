package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"projecttest/models"
)

func VerifyJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Lấy token từ header Authorization
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Không tìm thấy token")
		}

		// Giải mã token
		fmt.Println("=====token: ", tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra kiểu thuật toán
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(" không hợp lệ")
			}

			// Trả về secret key
			return []byte("mysecret"), nil
		})
		if err != nil {
			panic(err.Error())
			return c.JSON(http.StatusUnauthorized, "token khong hop le")
		}

		// Kiểm tra token
		if claims, ok := token.Claims.(*models.JwtCustomClaims); ok && token.Valid {
			c.Set("user", claims.Username)
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, "token khong hop le")
		}
	}
}
