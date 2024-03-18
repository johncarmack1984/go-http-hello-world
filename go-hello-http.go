package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.RemoteAddr)
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

func main() {
	port := getPort()

	log.Println("Starting server on port ", port)

	http.HandleFunc("/", logging(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}))

	http.ListenAndServe(port, nil)
}
