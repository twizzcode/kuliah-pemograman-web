package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Agar di terminal terlihat ada yang mengakses
	fmt.Printf("Ada yang akses:", r.Method, r.URL.Path)

	// Tulis response ke browser
	fmt.Fprint(w, "Hello! Anda berhasil akses web server GO!")
}

// Printf dan Println itu di terminal
// kalau Fprint itu di website

func main() {

	http.HandleFunc("/", helloHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	fmt.Println("Tekan Ctrl+C untuk menghentikan")

	http.ListenAndServe(":8080", nil)
}