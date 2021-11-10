package config

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

type SqlMock struct {
	Db   *sql.DB
	Mock sqlmock.Sqlmock
}

func InitDBMock() (sqlMock *SqlMock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("failed to create db mock: %s", err)
		return
	}
	sqlMock = &SqlMock{
		Db:   db,
		Mock: mock,
	}
	return
}
