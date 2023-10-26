// main.go
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	targetURL, _ := url.Parse("https://jsonplaceholder.typicode.com")

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/todos/1" {
			// forwarding request
			proxy.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
