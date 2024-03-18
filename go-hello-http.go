package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.Method, r.URL.Path)
		f(w, r)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":80"
	} else {
		port = ":" + port
	}

	return port
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func main() {
	port := getPort()

	log.Println("Starting server on port ", port)

	http.HandleFunc("/", logging(hello))

	http.HandleFunc("/health", logging(healthCheck))

	http.ListenAndServe(port, nil)
}
