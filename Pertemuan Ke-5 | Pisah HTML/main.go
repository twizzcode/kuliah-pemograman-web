package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type User struct {
	Nama  string
	Email string
}

type Produk struct {
	Nama  string
	Harga int
}

var DataProduk []Produk

var DataUser []User

var tplAbout = template.Must(template.ParseFiles(
	"template/layout.html",
	"template/about.html",
))

var tplProduk = template.Must(template.ParseFiles(
	"template/layout.html",
	"template/produk.html",
))

var tplProdukForm = template.Must(template.ParseFiles(
	"template/layout.html",
	"template/produk_form.html",
))

var tplHome = template.Must(template.ParseFiles(
	"template/layout.html",
	"template/home.html",
))

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /about", aboutHandler)
	mux.HandleFunc("GET /produk", produkListHandler)
	mux.HandleFunc("GET /produk/form", produkFormHandler)
	mux.HandleFunc("POST /produk/simpan", produkSimpanHandler)

	http.ListenAndServe(":8000", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// data := []string {
	// 	"ayam goreng", "bebek goreng", "nila goreng", "lele goreng",
	// }

	DataUser = append(DataUser, User{
		Nama:  "Budi 14",
		Email: "budi14@gmail.com",
	})
	DataUser = append(DataUser, User{
		Nama:  "Budi 15",
		Email: "budi15@gmail.com",
	})

	var data = map[string]any{
		"judul": "Halaman Home",
		"user":  DataUser,
	}

	tplHome.ExecuteTemplate(w, "layout", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tplAbout.ExecuteTemplate(w, "layout", nil)
}

func produkListHandler(w http.ResponseWriter, r *http.Request) {
	// siapkan data + nomor urut (mulai dari 1) untuk ditampilkan di tabel
	type barisProduk struct {
		No    int
		Nama  string
		Harga int
	}

	var daftar []barisProduk
	for i, p := range DataProduk {
		daftar = append(daftar, barisProduk{
			No:    i + 1,
			Nama:  p.Nama,
			Harga: p.Harga,
		})
	}

	data := map[string]any{
		"produk": daftar,
	}
	tplProduk.ExecuteTemplate(w, "layout", data)
}

func produkFormHandler(w http.ResponseWriter, r *http.Request) {
	tplProdukForm.ExecuteTemplate(w, "layout", nil)
}

func produkSimpanHandler(w http.ResponseWriter, r *http.Request) {
	nama := r.FormValue("nama")
	harga := r.FormValue("harga")

	hargaInt, _ := strconv.Atoi(harga)

	DataProduk = append(DataProduk, Produk{
		Nama:  nama,
		Harga: hargaInt,
	})

	http.Redirect(w, r, "/produk", http.StatusSeeOther)
}
