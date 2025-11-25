package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func Add(a int, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		a := 3
		b := 5
		result := Add(a, b)
		fmt.Fprintf(w, "The sum of %d and %d is %d", a, b, result)
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
