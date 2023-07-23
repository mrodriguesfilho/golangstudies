package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	port := ":5050"
	println("listening on ", port)
	log.Fatal(http.ListenAndServe(port, server))
}
