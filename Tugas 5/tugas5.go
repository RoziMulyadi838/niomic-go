package main

import "fmt"
import "math/rand"

type pelajar struct {
    nama string
    umur int
}

func main(){
    var n1,n2,n3 pelajar;
    n1.nama = "Ahmad"
    n2.nama = "Abi"
    n3.nama = "Zaka"
    
    // tambahkan umur pada masing masing variabel struct yaitu n1 , n2 dan n3 , berikan nilai pada umur secara random dengan mengugnakan fungsi rand.Intn(100) maksimal nilai random 100
    n1.umur = rand.Intn(100)
    n2.umur = rand.Intn(100)
    n3.umur = rand.Intn(100)
    
    fmt.Println(n1.nama, n1.umur)
    fmt.Println(n2.nama, n2.umur)
    fmt.Println(n3.nama, n3.umur)        
}