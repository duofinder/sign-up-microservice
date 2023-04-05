package utils

import "testing"

func Test_Utils_GenerateRefreshToken(t *testing.T) {
	_, err := GenerateRefreshToken()
	if err != nil {
		t.Fatal(err)
	}
}
