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
func DataPengunjung(){
	datamember := DataMember{
		Nama: "Damian Sitanggang",
		StatusMember: false,
		PinjamBuku: true,
		JudulBuku: "Buku Sains",
		LamaPeminjaman: 5,
	}
	fmt.Println(datamember)
}


// function
func Welcome() {
	fmt.Println("Perpustakaan Kota")
}

// function sebagai parameter
func Verifikasi(namaDepan string, namaBelakang string) {
	fmt.Println("Verifikasi", namaDepan, namaBelakang, "Selamat Datang Di Perpustakaan")
}
// Struct Method
func (member DataMember) StructMethod() {
	fmt.Println(member.Nama)
}

// funtion dengan return value
func Dendabuku(hargaDenda, LamaPeminjaman int) int{
	denda := hargaDenda * LamaPeminjaman
	return denda
	}

// function multiple return value
func PowerCal(I int, V int) (int, int) {
	Power := V * I
	Resistor := V % I
	return Power, Resistor
}
// Anonymous Struct
func AnonymousDataBuku() {
	DataBuku := struct {
	Buku string
	Judul string
	StatusBuku bool
	}{
		Buku: "Novel",
		Judul: "Olympus",
		StatusBuku: true,
	}
	fmt.Println(DataBuku)	
}
