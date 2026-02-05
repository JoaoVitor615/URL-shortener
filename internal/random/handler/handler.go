package handler

import (
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/random/service"
)

type URLRandomHandler struct {
	service *service.URLRandomService
}

func NewURLRandomHandler(service *service.URLRandomService) *URLRandomHandler {
	return &URLRandomHandler{
		service: service,
	}
}

func (h *URLRandomHandler) Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test"))
}
