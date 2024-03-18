package main

import (
	"fmt"
	"net/http"
	"os"
)

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(getPort(), nil)
}
