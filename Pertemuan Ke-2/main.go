package main

import (
	"errors"
	"fmt"
)

type Item struct {
	nama	string
	harga	int64
	jumlah	int64
}

func main() {
	daftarBelanja := []Item {}

	telur := Item{
		nama:	"telur",
		harga:	0,
		jumlah:	2,
	}
	beras := Item{
		nama:	"beras",
		harga:	15000,
		jumlah:	1,
	}	

	daftarBelanja, err := tambah(daftarBelanja, telur)
	if err != nil {
		fmt.Println(err)
		return
	}

	daftarBelanja, err = tambah(daftarBelanja, beras)
	if err != nil {
		fmt.Println(err)
		return
	}

	tampilkan(daftarBelanja)

	total := total(daftarBelanja)

	fmt.Printf("\nTotal Belanja: Rp %d\n", total)

}

func tambah(daftarBelanja []Item, itemBaru Item) ([]Item, error) {
	if itemBaru.harga <= 0 {
		return daftarBelanja, errors.New("Harga harus lebih dari 0")
	}
	if itemBaru.jumlah <= 0 {
		return daftarBelanja, errors.New("Jumlah harus lebih dari 0")
	}

	daftarBelanja = append(daftarBelanja, itemBaru)

	return daftarBelanja, nil
}

func tampilkan(daftarBelanja []Item) {
	fmt.Println("====== Daftar Belanja Kamu ======")

	for i, item := range daftarBelanja {
		subtotal := item.harga * item.jumlah

		fmt.Printf(
			"%d. %s - Rp %d x %d = Rp %d\n",
			i + 1,
			item.nama,
			item.harga,
			item.jumlah,
			subtotal,
		)
	}
}

func total(daftarBelanja []Item) int64 {
	var total int64 = 0
	
	for _, item := range daftarBelanja {
		total += item.harga * item.jumlah
	}

	return total
}