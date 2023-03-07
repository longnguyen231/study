package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"projecttest/connect"
	models "projecttest/models"
	"strings"
)

func VerifyJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Lấy token từ header Authorization
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Không tìm thấy token")
		}

		// Giải mã token
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra kiểu thuật toán
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(" không hợp lệ")
			}

			// Trả về secret key
			return []byte("mysecretkey"), nil
		})
		if err != nil {
			log.Print(err.Error())
			log.Print(35)
			return c.JSON(http.StatusUnauthorized, "token khong hop le")
		}

		// Kiểm tra token
		claims, ok := token.Claims.(*models.JwtCustomClaims)
		log.Print("eroo", token.Claims)

		if ok && token.Valid {
			db := connect.GetDB()
			var user = &models.User{}
			err = db.QueryRow("select id,name from users where username =?", claims.Username).Scan(&user.Id, &user.Name)
			if err != nil {
				log.Print(98)
				log.Print(err.Error())
			}
			c.Set("user", user)
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, "token khong hop le")
		}
	}
}
