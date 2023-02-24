package connect

import (
	"database/sql"
	"log"
)

type DB struct {
	db *sql.DB
}

var db *DB

func (d *DB) New() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Print(3)
		panic(err.Error())
	}
	d.db = db
}

func GetDB() *sql.DB {
	return db.db
}
