package handler

import (
	"BookShop/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	service *service.BookService
}

func InitBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) BookRoutes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.getBooks)

	return r
}

func (h *BookHandler) getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}

}
