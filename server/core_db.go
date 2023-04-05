package main

import (
	"database/sql"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB or Collection
var (
	db_main     *gorm.DB
	db_main_sql *sql.DB
)

func startDBConnection() {

	logrus.Printf("[core][func: startDBConnection] Starting DB Connection...")

	initDBMain()

}

func closeDBConnections() {

	logrus.Printf("[core][func: closeDBConnections] Closing DB Connection...")

	closeDBMain()

}

func initDBMain() {

	logrus.Printf("[core][func: initDBMain] DB Main - Connecting %s", config.DSN)

	var err error

	db_main, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("[core][func: initDBMain] Failed when connecting to DB Main: %v", err)
		os.Exit(1)
		return
	}

	db_main_sql, err = db_main.DB()
	if err != nil {
		logrus.Fatalf("[core][func: initDBMain] Failed when initiate connection to DB Main: %v", err)
		os.Exit(1)
		return
	}

	db_main_sql.SetMaxIdleConns(0)
	db_main_sql.SetMaxOpenConns(100)

	err = db_main_sql.Ping()
	if err != nil {
		logrus.Fatalf("[core][func: initDBMain] Unable to Ping DB Main: %v", err)
		os.Exit(1)
		return
	}

	if getEnv("ENV", "LOCAL") == "LOCAL" {
		db_main = db_main.Debug()
	}

	logrus.Printf("[core][func: initDBMain] DB Main - Connected")

}

func closeDBMain() {

	logrus.Print("[core][func: closeDBMain] DB Main - Closing")

	var err error

	err = db_main_sql.Close()
	if err != nil {
		logrus.Fatalf("[core][func: closeDBMain] Failed when closing DB Main connection: %v", err)
		return
	}

	logrus.Println("[core][func: closeDBMain] DB Main - Closed")

}
