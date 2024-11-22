package main

import (
	"authentication/endpoints"
	"authentication/validations"
	"fmt"
	"log"
	"net/http"
)

const port string = ":8080"

func doMappings() {
	http.HandleFunc("/", endpoints.Home)
	http.HandleFunc("/login", endpoints.Login)
}

func StartServer() {
	mux := http.NewServeMux()
	handler := http.HandlerFunc(validations.ValidateRequest)
	mux.HandleFunc(handler)
	// Spinning up the server.
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	doMappings()

	log.Println("Started on port", port)
	fmt.Println("To close connection CTRL+C :-)")

	StartServer()
}
