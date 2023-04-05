package utils

import (
	"testing"
)

func Test_Utils_GenerateAccessToken(t *testing.T) {
	_, err := GenerateAccessToken(999666333000)
	if err != nil {
		t.Fatal(err)
	}
}
