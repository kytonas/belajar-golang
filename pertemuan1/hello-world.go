package main

import "fmt"

func main() {
	var nama string 

	fmt.Println("Masukkan nama anda dengan benar")
	fmt.Scanln(&nama)

	switch nama {
	case "Ksatria": 
		fmt.Println("Namanya benar Ksatria")
	case "Nasywaan":
		fmt.Println("Namanya benar Nasywaan")
	default:
		fmt.Println("Namanya tidak terdaftar")
	}
}


