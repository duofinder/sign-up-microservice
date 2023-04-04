package repositories

import (
	"database/sql"
	"fmt"
)

func CreateAuthRepository(db *sql.DB, contact, passwordHash, refreshToken string) error {
	sqmt, err := db.Prepare(`INSERT INTO auths("contact", "password", "refresh_token") VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}

	result, err := sqmt.Exec(contact, passwordHash, refreshToken)
	if err != nil {
		return err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAff != 1 {
		return fmt.Errorf("Internal Server Error")
	}

	return nil
}
