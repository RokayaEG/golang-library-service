package database

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQLStorage(cfg mysql.Config) (*sqlx.DB, error) {

	// dsn := "root:root@(localhost:3306)/librarydb"
	db, err := sqlx.Connect("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func InitStorage(db *sqlx.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
