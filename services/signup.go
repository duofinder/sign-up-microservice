package services

import (
	"net/http"

	"github.com/duofinder/sign-up-microservice/types"
)

func SignupService(input *types.SignupInput) *types.Response {
	hash, err := input.EncryptPasswordFunc(input.Password)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       types.Map{"error": err.Error()},
		}
	}

	refreshToken, err := input.GenerateRefreshTokenFunc()
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       types.Map{"error": err.Error()},
		}
	}

	userId, err := input.CreateAuthRepository(input.DB, input.Contact, hash, refreshToken)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       types.Map{"error": err.Error()},
		}
	}

	accessToken, err := input.GenerateAccessTokenFunc(userId)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       types.Map{"error": err.Error()},
		}
	}

	return &types.Response{
		StatusCode: http.StatusCreated,
		Body: types.Map{
			"userId":       userId,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	}
}
