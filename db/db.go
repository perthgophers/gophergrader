package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func Init(dbusername, dbkey string) {
	var err error
	db, err = sqlx.Open("postgres", fmt.Sprintf("dbname=govhack user=%s password=%s host=103.241.200.20 sslmode=disable", dbusername, dbkey))

	if err != nil {
		log.Fatal(err)
	}
}

func Client() *sqlx.DB {
	return db
}
