package server

import (
	numeric_handler "github.com/JoaoVitor615/URL-shortener/internal/numeric/handler"
	numeric_service "github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
	random_handler "github.com/JoaoVitor615/URL-shortener/internal/random/handler"
	random_service "github.com/JoaoVitor615/URL-shortener/internal/random/service"
)

type Dependencies struct {
	NumericHandler   numeric_handler.INumericHandler
	URLRandomHandler *random_handler.URLRandomHandler
}

func NewDependencies() *Dependencies {
	numericService := numeric_service.NewNumericService()
	randomService := random_service.NewURLRandomService()
	return &Dependencies{
		NumericHandler:   numeric_handler.NewNumericHandler(numericService),
		URLRandomHandler: random_handler.NewURLRandomHandler(randomService),
	}
}
