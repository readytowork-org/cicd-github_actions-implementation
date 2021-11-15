package main

import (
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", helloHandler)
	http.HandleFunc("/", GetLocalIP)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Tada, 👻 Server is up and running!\n")
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP(w http.ResponseWriter, req *http.Request) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				io.WriteString(w, "Your IP address is\n")
			}
		}
	}
}
