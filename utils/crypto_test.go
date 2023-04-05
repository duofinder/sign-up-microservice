package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_Utils_Crypto_EncryptPassword(t *testing.T) {
	hash, err := EncryptPassword("123")
	if err != nil {
		t.Error(err)
	}

	if hash == "" {
		t.Error("Hash should not be a empty string")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte("123"))
	if err != nil {
		t.Error(err)
	}
}
