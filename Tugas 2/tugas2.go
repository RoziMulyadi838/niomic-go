package main

import "fmt"

func main() {
	var angka = 1

	if angka%2 == 0 {
		fmt.Println("Bilangan Genap")
	} else {
		fmt.Println("Bilangan Ganjil")
	}

	/* menggunakan switch case

	   switch angka % 2 {
	       case 0:
	           fmt.Println("Bilangan Genap")
	       default:
	           fmt.Println(" Bilangan Ganjil")
	   }
	*/
}