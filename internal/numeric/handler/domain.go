package handler

import (
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
)

type INumericHandler interface {
	GetLongURL(w http.ResponseWriter, r *http.Request)
	CreateShortURL(w http.ResponseWriter, r *http.Request)
}

func NewNumericHandler(service service.INumericService) INumericHandler {
	return &NumericHandler{
		service: service,
	}
}
