package main

import (
	"fmt"
	"net/http"
)

func main () {
	mux := http.NewServeMux()

	// route -> handler function
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /about", aboutHandler)
	mux.HandleFunc("GET /product/{id...}", productHandler)
	mux.HandleFunc("GET /product/test", productTestHandler)


	http.ListenAndServe(":8080", mux)
	
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is homePage")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is aboutPage")
}

func productTestHandler(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query().Get("mode")
	p := r.URL.Query().Get("page")
	full_url := r.URL.Path
	
	fmt.Fprint(w, "This is productTestPage\n")
	fmt.Fprintln(w, "Mode = ", m)
	fmt.Fprintln(w, "Page = ", p)
	fmt.Fprintln(w, "Full Path = ", full_url)
		
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprint(w, "This is productPage ", id)
}