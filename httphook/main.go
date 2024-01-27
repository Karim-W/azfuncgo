package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	port, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
