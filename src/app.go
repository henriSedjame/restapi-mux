package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/hsedjame/restapi-mux/src/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
}

func (app App) Start() {

	logger := log.New(os.Stdout, "[[io.github.hsejame.restapi-mux]] ", log.LstdFlags)

	router := mux.Router{}

	api.Controller{}.AddRoutes(&router)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           &router,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Second,
	}

	runAndStopServerGracefully(logger, server)

}

func runAndStopServerGracefully(logger *log.Logger, server *http.Server) {
	go func() {
		logger.Fatal(server.ListenAndServe())
	}()

	// Créer un chanel qui reçoit des signaux de type os.Signal
	signalChannel := make(chan os.Signal)

	// Envoyer un signal au channel lors :
	// * d'une interruption de la machine
	// * d'un arrêt de la machine
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// Récupérer le signal envoyé

	interruptOrKillSignal := <-signalChannel

	logger.Println("Interruption du server ==> ", interruptOrKillSignal)

	deadline, _ := context.WithTimeout(context.Background(), 30*time.Second)

	logger.Fatal(server.Shutdown(deadline))
}

func main() {
	App{}.Start()
}
