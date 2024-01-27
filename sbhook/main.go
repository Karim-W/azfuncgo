package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func queueHandler(w http.ResponseWriter, r *http.Request) {
	byts := make([]byte, r.ContentLength)
	r.Body.Read(byts)

	fmt.Println("Request Body: ", string(byts))
}

func main() {
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		customHandlerPort = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, func() http.Handler {
		return http.HandlerFunc(queueHandler)
	}()))
}
