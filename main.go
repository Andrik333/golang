package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.GetHandler)

	err := http.ListenAndServe(":8282", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
