package utils

import (
	"testing"
)

func Test_Utils_Crypto_EncryptPassword(t *testing.T) {
	hash, err := EncryptPassword("123")
	if err != nil {
		t.Error(err)
	}

	if hash == "" {
		t.Error("Hash should not be a empty string")
	}
}

func Test_Utils_Crypto_ComparePasswordWithHash(t *testing.T) {
	hash, err := EncryptPassword("123")
	if err != nil {
		t.Error(err)
	}

	ok := ComparePasswordWithHash("123", hash)
	if !ok {
		t.Error("The password does not match")
	}
}
