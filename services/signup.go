package services

import (
	"net/http"

	"github.com/duofinder/sign-up-microservice/types"
	"github.com/gin-gonic/gin"
)

func SignupService(input *types.SignupInput) *types.Response {
	hash, err := input.EncryptPasswordFunc(input.Password)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       gin.H{"error": err.Error()},
		}
	}

	refreshToken, err := input.GenerateRefreshTokenFunc()
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       gin.H{"error": err.Error()},
		}
	}

	userId, err := input.CreateAuthRepository(input.DB, input.Contact, hash, refreshToken)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       gin.H{"error": err.Error()},
		}
	}

	accessToken, err := input.GenerateAccessTokenFunc(userId)
	if err != nil {
		return &types.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       gin.H{"error": err.Error()},
		}
	}

	return &types.Response{
		StatusCode: http.StatusCreated,
		Body: gin.H{
			"userId":       userId,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	}
}
