package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Foo       string `json:"foo,omitempty"`
}

func main() {
	targetHost, _ := url.Parse("https://jsonplaceholder.typicode.com")

	//fastest way to create a reverse proxy is to use this function
	proxy := httputil.NewSingleHostReverseProxy(targetHost)

	//handle requires a path and a Handler. HandlerFunc lets us configure the Handler easier as Handler requires an object
	http.Handle("/todos/1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, targetHost, proxy)
	}))
	//we start the server and use log.Fatal like in the documentation
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request, target *url.URL, proxy *httputil.ReverseProxy) {
	//get target host
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}
	//in case of error exit function
	defer response.Body.Close()

	//where we store the data
	todo := Todo{}

	//NewDecoder reads the response and Decode fills in the fields
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data from Api: %+v", todo)
	//modify data
	todo.Foo = "bar"

	//encode todo to json format
	json.NewEncoder(w).Encode(todo)
}
