package types

import "database/sql"

type SignupInput struct {
	UserData
	DB                       *sql.DB
	EncryptPasswordFunc      func(password string) (string, error)
	GenerateRefreshTokenFunc func() (string, error)
	CreateAuthRepository     func(db *sql.DB, contact, passwordHash, refreshToken string) (int64, error)
	GenerateAccessTokenFunc  func(userId int64) (string, error)
}
