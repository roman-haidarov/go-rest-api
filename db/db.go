package db

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cnf mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cnf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
