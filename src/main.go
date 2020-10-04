package main

import (
	"github.com/gorilla/mux"
	"github.com/hsedjame/restapi-mux/src/handlers"
	"github.com/hsedjame/restapi-mux/src/repository"
	"log"
	"net/http"
)

func main() {

	routeHandlers := handlers.NewRouteHandlers(repository.NewBookRepositoryImpl())

	//Init a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/books", routeHandlers.GetBooks).Methods("GET")

	router.HandleFunc("/api/books/{id}", routeHandlers.GetBook).Methods("GET")

	router.HandleFunc("/api/books", routeHandlers.CreateBook).Methods("POST")

	router.HandleFunc("/api/books", routeHandlers.UpdateBook).Methods("PUT")

	router.HandleFunc("/api/books/{id}", routeHandlers.DeleteBook).Methods("DELETE")

 	// Launch server
	log.Fatal(http.ListenAndServe(":8080", router))
}
