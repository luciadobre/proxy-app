package main

import (
	"log"
	"net/http"

	"github.com/luciadobre/proxy-app/server"
)

func main() {
	handler := http.HandlerFunc(server.Server)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
