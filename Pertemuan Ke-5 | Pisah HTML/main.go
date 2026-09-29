package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Nama	string
	Email	string
}

var DataUser []User

var tplAbout = template.Must(template.ParseFiles(
	"template/about.html",
))

var tplProduct = template.Must((template.ParseFiles(
	"template/layout.html",
	"template/product.html",
)))

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /about", aboutHandler)
	mux.HandleFunc("GET /product", productHandler)
	
	http.ListenAndServe(":8000", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	var tplHome, err = template.ParseFiles("template/home.html")
	if err != nil {
		
	}
	
	// data := []string {
	// 	"ayam goreng", "bebek goreng", "nila goreng", "lele goreng",
	// }
	

	DataUser = append(DataUser, User{
		Nama: "Budi 14",
		Email: "budi14@gmail.com",
	})
	DataUser = append(DataUser, User{
		Nama: "Budi 15",
		Email: "budi15@gmail.com",
	})

	var data = map[string]any{
		"judul":	"Halaman Home",
	 	"user":		DataUser,
	}
	
	tplHome.Execute(w, data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	tplAbout.ExecuteTemplate(w, "about", nil)
}

func productHandler(w http.ResponseWriter, r *http.Request){
	data := 89
	tplProduct.ExecuteTemplate(w, "layout", data)
}