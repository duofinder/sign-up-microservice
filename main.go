package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/duofinder/auth-microservice/handlers"
	microserviceutils "github.com/duofinder/microservice-utils"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.New()

	logger := log.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	db, err := microserviceutils.NewPostgres(os.Getenv("DATABASE_POSTGRES_CONNSTR"))
	if err != nil {
		log.Fatal("error trying to create new server")
	}

	handlers.SignupRoute(router, db)

	srv, err := microserviceutils.NewServer(router, logger, os.Getenv("SERVER_ADDR"))
	if err != nil {
		log.Fatal("error trying to create new server")
	}

	err = microserviceutils.ListenAndServeWithGracefullyShutdown(srv, ctx, stop)
	if err != nil {
		log.Fatal(err)
	}
}
