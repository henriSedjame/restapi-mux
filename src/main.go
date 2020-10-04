package main

import (
	"github.com/hsedjame/restapi-mux/src/api"
	"log"
	"net/http"
)

func main() {

	// Launch server
	log.Fatal(http.ListenAndServe(":8080", &api.NewController().Router))
}
