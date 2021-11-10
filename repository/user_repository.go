package repository

import (
	"database/sql"
	"log"
	"time"
)

type UserRepository struct {
	UserId       string    `db:"user_id"`
	CrawlingDate time.Time `db:"crawling_date"`
}

func (u *UserRepository) create(db *sql.DB) (id int64, err error) {
	stmt, err := db.Prepare(userCreateStmt())
	if err != nil {
		log.Fatalf("failed to prepare user create statement: %s", err)
		return
	}

	result, err := stmt.Exec(u.UserId, u.CrawlingDate, u.CrawlingDate)
	if err != nil {
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
