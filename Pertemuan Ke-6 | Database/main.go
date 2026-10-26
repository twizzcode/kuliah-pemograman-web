package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type PageProduk struct {
	Title  string
	Produk []Produk
	Contoh Produk
}

type Produk struct {
	ID    int64  // BIGINT
	Name  string // VARCHAR
	Price int64  // BIGINT
	Stok  int    // INT
}

var DataProduk []Produk

var tplAbout = template.Must(template.ParseFiles(
	"template/layout.html",
	"template/about.html",
))

var tplProduk = template.Must(template.New("produk").Funcs(template.FuncMap{
	"add": func(a, b int) int { return a + b },
}).ParseFiles(
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN belum di-set (cek file .env)")
	}

	db, err := sql.Open("mysql", dsn) // db bertipe *sql.DB
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("database tidak bisa dihubungi: %v", err)
	}
	log.Println("database terhubung")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /profil", profilHandler)
	mux.HandleFunc("GET /produk", produkHandler(db))
	mux.HandleFunc("GET /produk/form", produkFormHandler)
	mux.HandleFunc("POST /produk", produkSimpanHandler(db))

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("server berjalan di http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplHome.ExecuteTemplate(w, "layout", nil)
}

func profilHandler(w http.ResponseWriter, r *http.Request) {
	tplAbout.ExecuteTemplate(w, "layout", nil)
}

// ambilSemuaProduk membaca semua baris dari tabel produk
func ambilSemuaProduk(ctx context.Context, db *sql.DB) ([]Produk, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, nama, harga, stok FROM produk ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []Produk
	for rows.Next() {
		var p Produk
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stok); err != nil {
			return nil, err
		}
		hasil = append(hasil, p)
	}
	return hasil, rows.Err()
}

// simpanProduk menambah satu baris ke tabel produk
func simpanProduk(ctx context.Context, db *sql.DB, p Produk) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO produk (nama, harga, stok) VALUES (?, ?, ?)",
		p.Name, p.Price, p.Stok,
	)
	return err
}

func produkHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		produk, err := ambilSemuaProduk(r.Context(), db)
		if err != nil {
			log.Println("ambil produk:", err)
			http.Error(w, "Gagal mengambil data produk", http.StatusInternalServerError)
			return
		}

		data := PageProduk{Title: "Produk", Produk: produk}
		tplProduk.ExecuteTemplate(w, "layout", data)
	}
}

func produkFormHandler(w http.ResponseWriter, r *http.Request) {
	tplProdukForm.ExecuteTemplate(w, "layout", nil)
}

func produkSimpanHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		price, err := strconv.ParseInt(r.FormValue("price"), 10, 64)
		if name == "" || err != nil {
			http.Error(w, "Nama dan harga wajib diisi dengan benar", http.StatusBadRequest)
			return
		}

		stok, err := strconv.Atoi(r.FormValue("stok"))
		if err != nil {
			stok = 0
		}

		if err := simpanProduk(r.Context(), db, Produk{Name: name, Price: price, Stok: stok}); err != nil {
			log.Println("simpan produk:", err)
			http.Error(w, "Gagal menyimpan produk", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/produk", http.StatusSeeOther)
	}
}
