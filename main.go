package main

import (
	"handlers"
	"log"
	"net/http"
)

func main() {
	r := handlers.NewRouter()

	err := http.ListenAndServe(":8282", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
