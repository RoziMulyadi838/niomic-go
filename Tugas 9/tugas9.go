package main

import "fmt"
import "strconv"
// import "os"

// func main() {
// 	defer fmt.Println("Hallo") // defer = menahan perintah ini untuk menjalankan kode dibaris berikutnya
// 	fmt.Println("Selamat Datang")
// 	fmt.Println("Selamat")
// 	os.Exit(1) // menghentikan baris perintah berikutnya dan tidak menampilkan baris yang diberi defer
// 	tampil()

// 	// output dari penggunaan defer :
// 	// Selamat Datang
// 	// Selamat
// 	// Tampilkan
// 	// Hallo

// 	// output dari penggunaan exit :
// 	// 	Selamat Datang
// 	// Selamat
// 	// exit status 1	

// }

// func tampil()  {
// 	fmt.Println("Tampilkan")
// }

func main() {
	// defer recover_saya()
	var input string
	fmt.Println("Masukkan Angka : ")
	fmt.Scanln(&input)

	var number int
	var err error

	number, err = strconv.Atoi(input)
	
	if err == nil {
		fmt.Println(number , "ini adalah angka")
	} else {
		fmt.Println(input, "Bukan Angka Ini")
		panic(err.Error())
		fmt.Println("Munculkan saya")
	}
}

func recover_saya() {
	if r:=recover(); r!=nil {
		fmt.Println("Akhirnya Saya Ditampilkan")
	} else {
		fmt.Println("Ini Lancar Saja")
	}
}