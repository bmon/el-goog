package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var databaseFile string = "elgoog.db"

var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", databaseFile)
	if err != nil {
		panic(err)
	}
}
