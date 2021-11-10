package repository

import (
	"database/sql"
	"log"
	"time"
)

type DetailRepository struct {
	BankName       string    `db:"bank_name"`
	TradingContent string    `db:"trading_content"`
	Amount         int64     `db:"amount"`
	UpdatedAt      time.Time `db:"updated_at"`
	UserId         int64     `db:"user_id"`
	BankId         int64     `db:"bank_id"`
}

func (d *DetailRepository) create(db *sql.DB) (id int64, err error) {
	stmt, err := db.Prepare(detailCreateStmt())
	if err != nil {
		log.Fatalf("failed to prepare detail create statement: %s", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.BankName, d.TradingContent, d.Amount, d.UpdatedAt, d.UserId, d.BankId)
	if err != nil {
		return
	}
	return
}
