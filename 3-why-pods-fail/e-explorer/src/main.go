package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var port int
	var err error
	portStr := os.Getenv("PORT")
	if portStr != "" {
		port, err = strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Cannot parse PORT " + portStr + "to int")
		}
		fmt.Println("Found env var PORT:", port)

	} else {
		log.Fatalf("Cannot find expected env var: PORT")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /")
		fmt.Fprintf(w, "Explorer")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /healthz")
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	})

	listen := ":" + strconv.Itoa(port)
	fmt.Println("Server listening on", listen)
	http.ListenAndServe(listen, nil)
}
