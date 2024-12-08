package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"github.com/go-sql-driver/mysql"
	"github.com/roman-haidarov/go-rest-api/cmd/api"
	"github.com/roman-haidarov/go-rest-api/config"
	"github.com/roman-haidarov/go-rest-api/db"
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

	go func() {
		if err := server.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error occurred while running server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	server.GracefulShutdown()
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
