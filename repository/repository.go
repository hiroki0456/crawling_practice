package repository

import "database/sql"

type Repository interface {
	create(db *sql.DB) (id int64, err error)
}

func Insert(db *sql.DB, r Repository) (id int64, err error) {
	id, err = r.create(db)
	if err != nil {
		return
	}
	return
}
