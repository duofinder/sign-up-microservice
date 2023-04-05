package repositories

import (
	"database/sql"
)

func CreateAuthRepository(db *sql.DB, contact, passwordHash, refreshToken string) (int64, error) {
	sqmt, err := db.Prepare(`INSERT INTO auths("contact", "password", "refresh_token") VALUES ($1, $2, $3) RETURNING id`)
	if err != nil {
		return 0, err
	}

	result := sqmt.QueryRow(contact, passwordHash, refreshToken)
	if err != nil {
		return 0, err
	}

	var userId int64

	if err = result.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}
