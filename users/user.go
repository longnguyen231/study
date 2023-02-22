package users

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"projecttest/connect"
	"projecttest/models"
)

func CreateUser(c echo.Context) error {

	u := new(models.User)

	if err := c.Bind(u); err != nil {
		log.Print(2)
		return err
	}

	db, err := connect.ConnectData()
	if err != nil {
		log.Print(23)
		panic(err.Error())
	}
	insert, err := db.Prepare("insert into users(Id,Name,Age,home_town) values (?,?,?,?)")
	if err != nil {
		log.Print(4)
		panic(err.Error())
	}
	_, err = insert.Exec(u.Id, u.Name, u.Age, u.Hometown)
	if err != nil {
		log.Print(5)
		panic(err.Error())
	}
	return c.JSON(http.StatusCreated, u)
}
func GetUser(c echo.Context) error {
	db, err := connect.ConnectData()
	if err != nil {
		log.Print(44)
		panic(err.Error())
	}
	var list = []models.User{}
	result, err := db.Query("select * from users  ")
	if err != nil {
		log.Fatal(err)

	}
	for result.Next() {
		var Id int
		var Name string
		var Age int
		var Hometown string
		var err = result.Scan(&Id, &Name, &Age, &Hometown)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, models.User{Id, Name, Age, Hometown})
	}
	return c.JSON(http.StatusOK, list)
}
func UpdateUser(c echo.Context) error {
	u := new(models.User)
	err := c.Bind(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	db, err := connect.ConnectData()
	if err != nil {
		log.Print(74)
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("update users set Name=?,Age=?,home_town=? where id=?", u.Name, u.Age, u.Hometown, u.Id)
	if err != nil {
		log.Print(79)
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, u)
}
func DeleteUser(c echo.Context) error {
	u := new(models.User)
	err := c.Bind(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	db, err := connect.ConnectData()

	if err != nil {
		log.Print(86)
		panic(err.Error())
	}
	_, err = db.Exec("DELETE FROM users WHERE id = ?", u.Id)
	if err != nil {
		log.Print(91)
		panic(err.Error())
	}
	return c.JSON(200, "delete succesful")
}
