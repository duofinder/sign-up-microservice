package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/duofinder/sign-up-microservice/repositories"
	"github.com/duofinder/sign-up-microservice/services"
	"github.com/duofinder/sign-up-microservice/types"
	"github.com/duofinder/sign-up-microservice/utils"
	"github.com/duofinder/sign-up-microservice/validation"

	_ "github.com/lib/pq"
)

func Signup(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST, OPTIONS",
		"Content-Type":                 "application/json",
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_POSTGRES_CONNSTR"))
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server Error",
			Headers:    headers,
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	defer db.Close()

	login, err := validation.Validate(request.Body, &types.UserData{})
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Bad Request. Body provided is incorrect!",
			Headers:    headers,
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	resp := services.SignupService(&types.SignupInput{
		UserData:                 login,
		DB:                       db,
		EncryptPasswordFunc:      utils.EncryptPassword,
		GenerateRefreshTokenFunc: utils.GenerateRefreshToken,
		CreateAuthRepository:     repositories.CreateAuthRepository,
		GenerateAccessTokenFunc:  utils.GenerateAccessToken,
	})

	body, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal Server Error",
			Headers:    headers,
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		Headers:    headers,
		StatusCode: resp.StatusCode,
	}, nil
}
