package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /")
		fmt.Fprintf(w, "Boundary")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /healthz")
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	})

	fmt.Println("Server listening on", port)
	http.ListenAndServe(port, nil)
}
