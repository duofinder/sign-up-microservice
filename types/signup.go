package types

import "database/sql"

type SignupInput struct {
	UserData
	DB                   *sql.DB
	EncryptPasswordFunc  func(password string) (string, error)
	CreateAuthRepository func(db *sql.DB, contact, passwordHash string) error
}
