package main

import (
	"log"
	"net/http"
)

func main() {
	mux := AttachHandler()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
