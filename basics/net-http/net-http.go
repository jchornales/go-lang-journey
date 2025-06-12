package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func helloHandler( w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, world!")
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func receiveJSON(w http.ResponseWriter, r *http.Request) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var p Person
	json.NewDecoder(r.Body).Decode(&p)
	fmt.Fprintf(w, "Received %s, age %d", p.Name, p.Age)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Incoming %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	finalHandler := http.HandlerFunc(helloHandler)
	http.Handle("/", loggingMiddleware(finalHandler))
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/receive", receiveJSON)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}