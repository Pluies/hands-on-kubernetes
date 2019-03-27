package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			fmt.Fprint(w, "page not found")
		} else {
			fmt.Println("Handling request for /")
			fmt.Fprintf(w, "Fern")
		}
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /healthz")
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	})

	fmt.Println("Server listening on", port)
	http.ListenAndServe(port, nil)
}
