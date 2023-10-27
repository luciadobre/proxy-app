package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// func NewProxy(rawUrl string) (*httputil.ReverseProxy, error) {
//     url, err := url.Parse(rawUrl)
//     if err != nil {
//         return nil, err
//     }
//     proxy := httputil.NewSingleHostReverseProxy(url)
//     return proxy, nil
// }

// func main() {
//     proxy, err := NewProxy("https://jsonplaceholder.typicode.com/todos/1/")
//     if err != nil {
//         panic(err)
//     }

//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         proxy.ServeHTTP(w, r)
//     })

//     log.Fatal(http.ListenAndServe(":8080", nil))
// }

// type Todo struct {
// 	UserID    int    `json:"userId"`
// 	ID        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// 	Foo       string `json:"foo,omitempty"`
// }

// func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
// 	url, err := url.Parse(targetHost)
// 	if err != nil {
// 		return nil, err
// 	}
// 	proxy := httputil.NewSingleHostReverseProxy(url)
// 	return proxy, nil
// }

// func ModifyJSONResponse(responseBody io.ReadCloser) (Todo, error) {
// 	bodyBytes, err := io.ReadAll(responseBody)
// 	if err != nil {
// 		return Todo{}, err
// 	}

// 	todoItem := Todo{}
// 	err = json.Unmarshal(bodyBytes, &todoItem)
// 	if err != nil {
// 		return Todo{}, err
// 	}

// 	todoItem.Foo = "bar"

// 	return todoItem, nil
// }

// func main() {

// 	proxy, err := NewProxy("https://jsonplaceholder.typicode.com/todos/1/")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	http.HandleFunc("/todos/1", func(w http.ResponseWriter, r *http.Request) {
// 		proxy.ServeHTTP(w, r)
// 	})
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func main() {

// 	targetURL, _ := url.Parse("https://jsonplaceholder.typicode.com")

// 	proxy := httputil.NewSingleHostReverseProxy(targetURL)

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path == "/todos/1" {
// 			r.URL.Path = "/todos/1"
// 			proxy.ServeHTTP(w, r)
// 		} else {
// 			http.NotFound(w, r)
// 		}
// 	})

// 	http.ListenAndServe(":8080", nil)
// }

// struct for response + additional foo: bar field
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

	json.NewEncoder(w).Encode(todo)
}
