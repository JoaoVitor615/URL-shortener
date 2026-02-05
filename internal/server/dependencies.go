package server

import (
	numeric_handler "github.com/JoaoVitor615/URL-shortener/internal/numeric/handler"
	numeric_service "github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
	random_handler "github.com/JoaoVitor615/URL-shortener/internal/random/handler"
	random_service "github.com/JoaoVitor615/URL-shortener/internal/random/service"
)

type Dependencies struct {
	URLNumericHandler *numeric_handler.URLNumericHandler
	URLRandomHandler  *random_handler.URLRandomHandler
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		URLNumericHandler: numeric_handler.NewURLNumericHandler(numeric_service.NewURLNumericService()),
		URLRandomHandler:  random_handler.NewURLRandomHandler(random_service.NewURLRandomService()),
	}
}
