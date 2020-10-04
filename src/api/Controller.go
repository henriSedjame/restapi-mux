package api

import (
	"github.com/gorilla/mux"
	"github.com/hsedjame/restapi-mux/src/handlers"
	"github.com/hsedjame/restapi-mux/src/repository"
)

type Controller struct {
	Router mux.Router
}

func NewController() *Controller {

	routeHandlers := handlers.NewRouteHandlers(repository.NewBookRepositoryImpl())

	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/books", routeHandlers.GetBooks).Methods("GET")

	router.HandleFunc("/api/books/{id}", routeHandlers.GetBook).Methods("GET")

	router.HandleFunc("/api/books", routeHandlers.CreateBook).Methods("POST")

	router.HandleFunc("/api/books", routeHandlers.UpdateBook).Methods("PUT")

	router.HandleFunc("/api/books/{id}", routeHandlers.DeleteBook).Methods("DELETE")

	return &Controller{*router}
}
