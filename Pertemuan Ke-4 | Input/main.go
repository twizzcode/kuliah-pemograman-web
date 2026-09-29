package main

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"
)

type DataPendaftar struct {
	Nama   string
	Alamat string
}

var data_storage []DataPendaftar

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", homeHandler)

	mux.HandleFunc("GET /form", formHandler)
	mux.HandleFunc("POST /simpan", simpanHandler)
	

	http.ListenAndServe(":8000", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	mode := r.URL.Query().Get("mode")
	fmt.Fprintln(w, "Query string : ", query)
	fmt.Fprintln(w, "Mode string : ", mode)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	pendaftar := DataPendaftar{
		Nama:   "",
		Alamat: "",
	}
	html := tampilkanForm(pendaftar, nil)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

func tampilkanForm(data_form DataPendaftar, error []string) string {
	html := `
		<h5>Form Pendaftaran</h5>
		<p style="color: red;">%s</p>
		<form action="/simpan" method="POST">
			<table>
				<tr>
					<td>Nama Lengkap</td>
					<td><input type="text" id="nama" name="nama" value="%s"></td>
				</tr>
				<tr>
					<td>Alamat</td>
					<td><textarea id="alamat" name="alamat">%s</textarea></td>
				</tr>
				<tr>
					<td></td>
					<td><input type="submit" value="Submit"></td>
				</tr>
			</table>
		</form>
		%s
	`

	var str_error string
	if len(error) > 0 {
		str_error = strings.Join(error, "<br>")
	}

	str_data := ""
	for _, v := range data_storage {
		str_data += fmt.Sprintf("<li>%s - %s</li>", v.Nama, v.Alamat)
	}

	return fmt.Sprintf(html, str_error, data_form.Nama, data_form.Alamat, str_data)
}

func simpanHandler(w http.ResponseWriter, r *http.Request) {
	nama := r.PostFormValue("nama")
	alamat := r.PostFormValue("alamat")
	error := []string{}

	is_nama_ada_angka := false

	for _,v := range nama {
		if(unicode.IsDigit(v)) {
			is_nama_ada_angka = true
			break
		}
	}
	if(is_nama_ada_angka) {
		error = append(error, "Nama tidak boleh mengandung angka")
	}
	if(utf8.RuneCountInString(nama) > 20) {
		error = append(error, "Nama tidak boleh lebih dari 20 karakter")
	}
	
	if len(error) > 0 {
		data_form := DataPendaftar{
			Nama:   nama,
			Alamat: alamat,
		}
		html := tampilkanForm(data_form, error)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, html)
		return
	}

	data_storage = append(data_storage, DataPendaftar{
		Nama:   nama,
		Alamat: alamat,
	})

	http.Redirect(w,r, "form", http.StatusSeeOther)
	
}