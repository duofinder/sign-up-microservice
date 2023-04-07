package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/duofinder/sign-up-microservice/types"
)

func Test_Service_Signup(t *testing.T) {
	var userId int64 = 1
	refreshToken := "refreshtoken"
	accessToken := "accesstoken"

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
					UserData: &types.UserData{
						Contact:  "email-example@gmail.com",
						Password: "theworstpasswordpossible",
					},
					DB: nil,
					EncryptPasswordFunc: func(password string) (string, error) {
						return "hashedpassword", nil
					},
					GenerateRefreshTokenFunc: func() (string, error) {
						return refreshToken, nil
					},
					CreateAuthRepository: func(db *sql.DB, contact, passwordHash, refreshToken string) (int64, error) {
						return userId, nil
					},
					GenerateAccessTokenFunc: func(userId int64) (string, error) {
						return accessToken, nil
					},
				},
			},
			want: &types.Response{
				StatusCode: http.StatusCreated,
				Body: types.Map{
					"accessToken":  accessToken,
					"refreshToken": refreshToken,
					"userId":       userId,
				},
			},
		},
		{
			name: "Should return an error if encrypt password fail",
			args: args{
				&types.SignupInput{
					UserData: &types.UserData{
						Contact:  "email-example@gmail.com",
						Password: "theworstpasswordpossible",
					},
					DB: nil,
					EncryptPasswordFunc: func(password string) (string, error) {
						return "", fmt.Errorf("something failed")
					},
					GenerateRefreshTokenFunc: func() (string, error) {
						return refreshToken, nil
					},
					CreateAuthRepository: func(db *sql.DB, contact, passwordHash, refreshToken string) (int64, error) {
						return userId, nil
					},
					GenerateAccessTokenFunc: func(userId int64) (string, error) {
						return accessToken, nil
					},
				},
			},
			want: &types.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       types.Map{"error": "something failed"},
			},
		},
		{
			name: "Should return an error if repository fail",
			args: args{
				&types.SignupInput{
					UserData: &types.UserData{
						Contact:  "email-example@gmail.com",
						Password: "theworstpasswordpossible",
					},
					DB: nil,
					EncryptPasswordFunc: func(password string) (string, error) {
						return "hashedpassword", nil
					},
					GenerateRefreshTokenFunc: func() (string, error) {
						return refreshToken, nil
					},
					CreateAuthRepository: func(db *sql.DB, contact, passwordHash, refreshToken string) (int64, error) {
						return 0, fmt.Errorf("something failed")
					},
					GenerateAccessTokenFunc: func(userId int64) (string, error) {
						return accessToken, nil
					},
				},
			},
			want: &types.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       types.Map{"error": "something failed"},
			},
		},
		{
			name: "Should return an error if generate access token fail",
			args: args{
				&types.SignupInput{
					UserData: &types.UserData{
						Contact:  "email-example@gmail.com",
						Password: "theworstpasswordpossible",
					},
					DB: nil,
					EncryptPasswordFunc: func(password string) (string, error) {
						return "hashedpassword", nil
					},
					GenerateRefreshTokenFunc: func() (string, error) {
						return refreshToken, nil
					},
					CreateAuthRepository: func(db *sql.DB, contact, passwordHash, refreshToken string) (int64, error) {
						return userId, nil
					},
					GenerateAccessTokenFunc: func(userId int64) (string, error) {
						return "", fmt.Errorf("something failed")
					},
				},
			},
			want: &types.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       types.Map{"error": "something failed"},
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
