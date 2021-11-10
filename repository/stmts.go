package repository

func userCreateStmt() (s string) {
	s = `INSERT INTO
			Users(
				user_id,
				crawling_date
			)
			VALUES(
				?,
				?
			)
			ON DUPLICATE KEY UPDATE
				crawling_date = ?`
	return s
}

func bankCreateStmt() (s string) {
	s = `INSERT INTO
			Banks(
				name,
				updated_at,
				user_id
			)
			VALUES(
				?,
				?,
				?
			)`
	return s
}

func detailCreateStmt() (s string) {
	s = `INSERT INTO
			Details(
				bank_name,
				trading_content,
				amount,
				updated_at,
				user_id,
				bank_id
			)
			VALUES(
				?,
				?,
				?,
				?,
				?,
				?
			)`
	return s
}
