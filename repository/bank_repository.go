package repository

import (
	"database/sql"
	"log"
	"time"
)

type BankRepository struct {
	Name      string    `db:"crawling_date"`
	UpdatedAt time.Time `db:"updated_at"`
	UserId    int64     `db:"user_id"`
}

func (b *BankRepository) create(db *sql.DB) (id int64, err error) {
	stmt, err := db.Prepare(bankCreateStmt())
	if err != nil {
		log.Fatalf("failed to prepare bank create statement: %s", err)
		return
	}

	result, err := stmt.Exec(b.Name, b.UpdatedAt, b.UserId)
	if err != nil {
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
