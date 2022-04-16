package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/time", getTime)

	fmt.Printf("Starting server on ")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatal(err)
	}
}
