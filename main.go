package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/health-check", helloHandler)

	log.Println("Listning for requests at /health-check")
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(port, nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Tada, 👻 Server is up and running!\n")
}
