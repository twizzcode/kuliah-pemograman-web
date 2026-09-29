package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	// 1. Halaman utama memakai {$} -> hanya cocok untuk path "/" persis
	mux.HandleFunc("GET /{$}", homeHandler)

	// 2. Route statis biasa
	mux.HandleFunc("GET /about", aboutHandler)

	// 3. Route dengan path parameter {id}
	mux.HandleFunc("GET /product/{id}", productHandler)

	// 4. Route dengan query string + nilai default
	mux.HandleFunc("GET /search", searchHandler)

	// 5. Route POST yang membalas 201 Created
	mux.HandleFunc("POST /product", createProductHandler)

	log.Println("Server berjalan di http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

// GET /{$} -> halaman utama
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Home Page")
}

// GET /about
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is About Page")
}

// GET /product/{id} -> path parameter
func productHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "This is Product Page dengan id = %s", id)
}

// GET /search?q=...&page=... -> query string dengan default
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Query string dengan nilai default
	q := query.Get("q")
	if q == "" {
		q = "semua" // nilai default
	}

	// Query string dengan default angka + konversi
	page := 1
	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			page = parsed
		}
	}

	fmt.Fprintf(w, "Search: q = %q, page = %d\n", q, page)
}

// POST /product -> membalas 201 Created
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var payload map[string]any
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil && err.Error() != "EOF" {
		http.Error(w, "Body JSON tidak valid", http.StatusBadRequest)
		return
	}

	// Set status 201 Created secara eksplisit
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Produk berhasil dibuat",
		"data":    payload,
	})
}
