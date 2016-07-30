package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("postgres", "dbname=govhack user=postgres password=iamagopher host=103.241.200.20 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
}

func Client() *sqlx.DB {
	return db
}
