package dataksesmodule

import "fmt"

// function
func Welcome() {
	fmt.Println("Selamat Datang di Perpustakaan")
}

// function sebagai parameter
func Verifikasi(namaDepan string, namaBelakang string) {
	fmt.Println("Verifikasi", namaDepan, namaBelakang, "Selamat Datang Di Perpustakaan")
}

// funtion return value
func Dendabuku(hargaDenda, jumlahHari int) int{
	denda := hargaDenda * jumlahHari
	return denda
	}


