package dataksesmodule

import "fmt"

// struct
type DataMember struct {
	Nama string
	PinjamBuku bool
	StatusMember bool
	JudulBuku string
	LamaPeminjaman int
}
// function
func Welcome() {
	fmt.Println("Perpustakaan Kota")
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

// function multiple return value
func PowerCal(I int, V int) (int, int) {
	Power := V * I
	Resistor := V % I
	return Power, Resistor
}




