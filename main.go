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

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Tada, 👻 Server is up and running!\n")
}
