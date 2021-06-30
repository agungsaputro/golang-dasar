package main

import "fmt"

func main()  {
	var name string
	
	

	name="agung dwi"
	fmt.Println(name)

	name="agung saputro"
	fmt.Println(name)

	var aku = "agung" // variable langsung diisi data maka golang akan sudah mengindetifikasi type data
	fmt.Println(aku)

	var nilaiku = 12
	fmt.Println(nilaiku)

	tanggal := 15 // var bisa diganti := di deklarasi awal
	fmt.Println(tanggal)

	var(
		firstName = "Agung" // multiple deklarasi
		lastName = "Saputro"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}