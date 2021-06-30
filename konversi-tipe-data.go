package main

import (
	"fmt"
)

func main()  {
	var(
		nilai32 int32 = 1000
		nilai64 int64 = int64(nilai32)
		nilai8 int8 = int8(nilai32)
		name = "agung"
		e = name[0]
		estring string = string(e)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println(estring)
	
}