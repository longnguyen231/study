package connect

import (
	"database/sql"
	"log"
)

type DBSql struct {
	db *sql.DB
}

var dbSql *sql.DB

func (d *DBSql) New() {
	dbConnect, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Print(3)
		panic(err.Error())
	}
	dbSql = dbConnect
}

func GetDB() *sql.DB {
	return dbSql
}
