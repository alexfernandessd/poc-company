package main

import (
	"fmt"
	"log"
	"net/http"
	"poc-company/company"

	"github.com/facebookgo/grace/gracehttp"
	"go.uber.org/zap"
)

func main() {

	config := company.NewConfig()

	db, derr := company.NewDynamoDatabase(config.AWSRegion, config.DynamodbEndpoint, config.AWSMetadataURL)
	if derr != nil {
		log.Fatal("could not start service", zap.Error(derr))
	}

	repository := company.NewRepository(db)

	service := company.NewService(repository)

	handler := createHandler(service)

	fmt.Println("Starting server on port: ", config.Port)

	err := gracehttp.Serve(&http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: handler,
	})

	if err != nil {
		log.Fatal("failed on server start", err)
	}
}
