package main

import (
	"log"
	"os"

	mySqlConf "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/roman-haidarov/go-rest-api/config"
	"github.com/roman-haidarov/go-rest-api/db"
)

func main() {
	db, err := db.NewMySQLStorage(mySqlConf.Config{
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

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args) - 1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err == migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err == migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
