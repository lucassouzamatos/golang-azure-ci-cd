package main

import (
	"net/http"
	"log"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World by go handler"))
}

func main() {
	customHandlerPort, _ := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")

	mux := http.NewServeMux()
	mux.HandleFunc("/api/helloWorld", helloWorld)

	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}