package utils

import (
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
	 _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "tbilous"
    password = "bilous"
    dbname   = "tbilous"
)

func GetConnection() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	log.Println("DB Connection established...")
	return db
}
