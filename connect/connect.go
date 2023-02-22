package connect

import (
	"database/sql"
	"log"
)

func ConnectData() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Print(3)
		panic(err.Error())

	}
	return db, nil
}
