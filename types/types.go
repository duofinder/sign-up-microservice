package types

import "database/sql"

type Map map[string]any

type UserData struct {
	Contact  string `json:"contact" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=24"`
}

type Response struct {
	StatusCode int
	Body       Map
}

type SignupInput struct {
	*UserData
	DB                       *sql.DB
	EncryptPasswordFunc      func(password string) (string, error)
	GenerateRefreshTokenFunc func() (string, error)
	CreateAuthRepository     func(db *sql.DB, contact, passwordHash, refreshToken string) (int64, error)
	GenerateAccessTokenFunc  func(userId int64) (string, error)
}
