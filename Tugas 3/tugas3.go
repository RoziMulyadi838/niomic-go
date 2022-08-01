package main

import "fmt"

func main(){
    var buah = []string{"apel","jeruk","melon","anggur"}
    buah = append(buah, "pepaya")
    fmt.Println("Jumlah Element : ", len(buah))
    fmt.Println("Isi Element : ", (buah))
    
    /* contoh perulangan cara 1
    for i := 0; i < len(buah); i++ {
        fmt.Println("Element Ke - :",i ,buah[i])
    }
    */
    
    /* contoh perulangan cara 2
    var i = 0
    for i < len(buah) {
        fmt.Println("Element Ke - :",i,buah[i])
        i++
    }
    */
    
    // contoh perulangan cara 3
    var i = 0
    for{
        fmt.Println("Elemen Ke - :", i,buah[i])
        i++
        if i == len(buah){
            break
        }
    }
}