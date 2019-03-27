package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	port := ":8080"

	for i := 0; i <= 25; i++ {
		log.Println("Initialising... " + strconv.Itoa(4*i) + "%")
		time.Sleep(250 * time.Millisecond)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /")
		fmt.Fprintf(w, "Grey")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /healthz")
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	})

	fmt.Println("Server listening on", port)
	http.ListenAndServe(port, nil)
}
