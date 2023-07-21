package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	handlers "github.com/mohammedimrankasab/go-rest-proto/handlers"
)

const webServerPort = "8099"

func main() {

	app := handlers.Config{}

	fmt.Println("Starting the API server on port", webServerPort)
	r := mux.NewRouter()
	r.HandleFunc("/echo", app.Hello).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + webServerPort,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
