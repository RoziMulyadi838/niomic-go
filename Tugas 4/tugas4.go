package main

import "fmt"

func main() {
	// 1. buatlah sebuah map dengan nama mahasiswa yang berisi 2 buah nilai tinggi badan tiap mahasiswa
	mahasiswa := map[string]int{"Aldo": 182, "Yosep": 178}
	tampil_mahasiswa(mahasiswa)
}

// 2. setelah itu buatlah sebuah fungsi bernama tampil_mahasiswa lalu kirim nilai map ke dalam fungsi
func tampil_mahasiswa(mahasiswa map[string]int) {

	// 3. setelah teman teman buat fungsi tampil di fungsi tampil_mahasiswa sehingga menampilkan output
	// hint : Nama key pada map gunakan nama mahasiswanya
	for nama, umur := range mahasiswa {
		fmt.Println(nama, "\t:", umur, "cm")
	}
}