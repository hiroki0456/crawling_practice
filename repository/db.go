package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (db *sql.DB) {
	var err error
	db, err = sql.Open("mysql", "root@/mf")
	if err != nil {
		log.Fatalf("failed to open mysql: %s", err)
		return
	}
	return
}
