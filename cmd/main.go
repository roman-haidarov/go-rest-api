package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/roman-haidarov/go-rest-api/cmd/api"
	"github.com/roman-haidarov/go-rest-api/config"
	"github.com/roman-haidarov/go-rest-api/db"
	"log"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:              config.Envs.DBUser,
		Passwd:            config.Envs.DBPasswd,
		Addr:              config.Envs.DBAddress,
		DBName:            config.Envs.DBName,
		Net:               "tcp",
		AllowOldPasswords: true,
		ParseTime:         true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
