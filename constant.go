package main

import "fmt"

func main()  {
	const firstName = "Agung" // constan tidak bisa diubah , var bisa diubah
	const lastName = "Saputro"
	const value = 1000 // constan tidak akan error walau tidak digunakan

	const(
		alamat = "rembang"
		nomor = 12
	)



	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(alamat)
	fmt.Println(nomor)
	
}