package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	port := ":8080"

	sliceItemLength := 10000
	size := 0
	maxSize := 768
	step := maxSize / 64
	heavySlice := make([]string, maxSize)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /")
		fmt.Fprintf(w, "Highland")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /healthz")
		if size < maxSize {
			for i := 0; i < step; i++ {
				heavySlice = append(heavySlice, strings.Repeat("⚖️", sliceItemLength))
			}
			size += step
		}
		fmt.Fprintf(w, "{\"heavySlice\": "+strconv.Itoa(size)+",\"status\": \"ok\"}")
	})

	fmt.Println("Server listening on", port)
	http.ListenAndServe(port, nil)
}
