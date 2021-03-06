package dbConnect

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:yxHie25@localhost:5432/rfid?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

func GetDBConnect() *sql.DB {
	return DB
}
