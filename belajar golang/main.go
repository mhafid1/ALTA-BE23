package main

import "fmt"

func main() {
//	fmt.Println("hello world")
//INPUT (angka) -> menerima input / masukan dari user ke dalam program

//menerima input ANGKA dari user
//var angka int 
//fmt.Scanln(&angka)

//fmt.Println("hasil inputnya adalah", angka)

//var nama string
//fmt.Scanln(&nama)
//fmt.Println("hasil input nama adalah", nama)

//var kondisi bool
//fmt.Println("isi variable kondisi", kondisi)

//var angka int
//fmt.Println("masukkan angka untuk di cek:")
//fmt.Scanln(&angka)

//if angka%2 == 0 { 
//	fmt.Println("genap")
//} else {
//	fmt.Println("Ganjil")
//}

var umur int 
fmt.Println("mauskkan umur untuk di cek :")
fmt.Scanln(&umur)

if umur > 20 {
	fmt.Println("sudah dewasa")
	if umur >25 {
		fmt.Println("sudah tuwir")
} else {
	fmt.Println("kurang umur")
	
	}
}

if umur > 17 {
	fmt.Println("sudah dapat ktp")
} else {
	fmt.Println("belum dapat ktp")
}

}




