package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/duofinder/sign-up-microservice/types"
	"github.com/gin-gonic/gin"
)

func Test_Service_Signup(t *testing.T) {
	type args struct {
		input *types.SignupInput
	}
	tests := []struct {
		name string
		args args
		want *types.Response
	}{
		{
			name: "Should return success",
			args: args{
				&types.SignupInput{
					UserData: types.UserData{
						Contact:  "email-example@gmail.com",
						Password: "theworstpasswordpossible",
					},
					DB: nil,
					EncryptPasswordFunc: func(password string) (string, error) {
						return "hey, a hashed password", nil
					},
					GenerateRefreshTokenFunc: func() (string, error) {
						return "hey, a refresh token", nil
					},
					CreateAuthRepository: func(db *sql.DB, contact, passwordHash, refreshToken string) error {
						return nil
					},
				},
			},
			want: &types.Response{
				StatusCode: http.StatusCreated,
				Body:       gin.H{"message": "success"},
			},
		},
		{
			name: "Should return an error if encrypt password fail",
			args: args{
				&types.SignupInput{
					UserData: types.UserData{
						Contact:  "email-example@gmail.com",
						Password: "theworstpasswordpossible",
					},
					DB: nil,
					EncryptPasswordFunc: func(password string) (string, error) {
						return "", fmt.Errorf("something failed")
					},
					GenerateRefreshTokenFunc: func() (string, error) {
						return "hey, a refresh token", nil
					},
					CreateAuthRepository: func(db *sql.DB, contact, passwordHash, refreshToken string) error {
						return nil
					},
				},
			},
			want: &types.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       gin.H{"error": fmt.Errorf("something failed")},
			},
		},
		{
			name: "Should return an error if repository fail",
			args: args{
				&types.SignupInput{
					UserData: types.UserData{
						Contact:  "email-example@gmail.com",
						Password: "theworstpasswordpossible",
					},
					DB: nil,
					EncryptPasswordFunc: func(password string) (string, error) {
						return "hey, a hashed password", nil
					},
					GenerateRefreshTokenFunc: func() (string, error) {
						return "hey, a refresh token", nil
					},
					CreateAuthRepository: func(db *sql.DB, contact, passwordHash, refreshToken string) error {
						return fmt.Errorf("something failed")
					},
				},
			},
			want: &types.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       gin.H{"error": fmt.Errorf("something failed")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SignupService(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignupService() = %v, want %v", got, tt.want)
			}
		})
	}
}
