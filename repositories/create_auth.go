package repositories

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateAuthRepository(db *sql.DB, contact, passwordHash string) error {
	result, err := db.Exec(`INSERT INTO auths("contact", "password") VALUES ('$1', '$2')`, contact, passwordHash)
	if err != nil {
		log.Println(2)
		return err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		log.Println(3)
		return err
	}

	if rowsAff != 1 {
		log.Println(4)
		return fmt.Errorf("Internal Server Error")
	}

	return nil
}
