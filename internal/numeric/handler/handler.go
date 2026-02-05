package handler

import (
	"net/http"

	"github.com/JoaoVitor615/URL-shortener/internal/numeric/service"
)

type URLNumericHandler struct {
	service *service.URLNumericService
}

func NewURLNumericHandler(service *service.URLNumericService) *URLNumericHandler {
	return &URLNumericHandler{
		service: service,
	}
}

func (h *URLNumericHandler) Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test"))
}
