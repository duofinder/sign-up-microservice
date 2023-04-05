package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 39)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	s := base64.URLEncoding.EncodeToString(b)

	return s, nil
}
