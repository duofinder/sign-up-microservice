package repositories

import (
	"database/sql"
	"fmt"
)

func (rp *Repo) CreateAuthRepository(db *sql.DB, contact, passwordHash string) error {
	stmt, err := db.Prepare(`INSERT INTO auths("contact", "password") VALUES ('$1', '$2')`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(contact, passwordHash)
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
