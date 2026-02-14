package server

import (
	"github.com/JoaoVitor615/URL-shortener/internal/adapters"
	numeric_handler "github.com/JoaoVitor615/URL-shortener/internal/numeric/handler"
	"github.com/JoaoVitor615/URL-shortener/internal/numeric/repository"
	numeric_service "github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
	random_handler "github.com/JoaoVitor615/URL-shortener/internal/random/handler"
	random_service "github.com/JoaoVitor615/URL-shortener/internal/random/service"
)

type Dependencies struct {
	NumericHandler   numeric_handler.INumericHandler
	URLRandomHandler *random_handler.URLRandomHandler
}

func NewDependencies() *Dependencies {
	dynamoClient := adapters.InitializeDynamoClient()

	repository := repository.NewDynamoRepository(dynamoClient, "shortener-numeric")
	numericService := numeric_service.NewNumericService(repository)

	randomService := random_service.NewURLRandomService()
	return &Dependencies{
		NumericHandler:   numeric_handler.NewNumericHandler(numericService),
		URLRandomHandler: random_handler.NewURLRandomHandler(randomService),
	}
}
