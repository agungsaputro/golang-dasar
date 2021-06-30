package main

import "fmt"

func main()  {
	var a = 12
	var b = 12
	var c = 12 + 12
	var d = a + b
	var i = 10
	i += 10 //augmanted
	i++ //ternary operator

	var name1 = "agung"
	var name2 = "agung"
	var nilai1 = 12
	var nilai2 = 14

	var result bool = name1 == name2
	var result2 bool =  nilai1 > nilai2

	var jujur = true
	var bohong = false

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(i)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(nilai1 > nilai2)
	fmt.Println(nilai1 < nilai2)
	fmt.Println(nilai1 >= nilai2)
	fmt.Println(nilai1 <= nilai2)
	fmt.Println(nilai1 != nilai2)
	fmt.Println(jujur && bohong)
	fmt.Println(jujur || bohong)
}