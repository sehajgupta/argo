package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

var DB *sqlx.DB

func Init() {
	var err error
	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")

	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost)
    DB, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
}
