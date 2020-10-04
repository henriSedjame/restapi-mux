package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hsedjame/restapi-mux/src/models"
	"github.com/hsedjame/restapi-mux/src/repository"
	"net/http"
)

type RouteHandlers struct {
	BookRepo repository.BookRepository
}

func NewRouteHandlers(bookRepo repository.BookRepository) *RouteHandlers {
	return &RouteHandlers{BookRepo: bookRepo}
}

func (handler RouteHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(handler.BookRepo.GetBooks())
}

func (handler RouteHandlers) GetBook(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(handler.BookRepo.GetBook(id))
}

func (handler RouteHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(handler.BookRepo.CreateBook(book))
}

func (handler RouteHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(handler.BookRepo.UpdateBook(book))
}

func (handler RouteHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(handler.BookRepo.DeleteBook(id))
}
