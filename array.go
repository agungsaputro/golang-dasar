package main

import "fmt"

func main()  {
	var name = [3]string{
		"agung",
		"dwi",
		"saputro",
	}

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	fmt.Println(name)
	fmt.Println(len(name)) //panjang array
}