package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var pemain1, pemain2 string
	var pilihanPemain1, pilihanPemain2 int

	fmt.Println("Pilih mau pilihan apa ?")
	fmt.Println("Ketik 1 untuk gunting")
	fmt.Println("Ketik 2 untuk kertas")
	fmt.Println("Ketik 3 untuk batu")

	fmt.Scanln(&pilihanPemain1)

	switch pilihanPemain1 {
	case 1:
		pemain1 = "Gunting"
	case 2:
		pemain1 = "Kertas"
	case 3:
		pemain1 = "Batu"
	default:
		fmt.Println("Tidak ada pilihan")
	}

	pilihanPemain2 = rand.Intn(3)

	switch pilihanPemain2 {
	case 1:
		pemain2 = "Gunting"
	case 2:
		pemain2 = "Kertas"
	case 3:
		pemain2 = "Batu"
	default:
		fmt.Println("Tidak ada pilihan")
	}

	fmt.Printf("Pilihan pemain satu %s \n", pemain1)
	fmt.Printf("Pilihan pemain dua %s \n", pemain2)

	if pemain1 == "Kertas" && pemain2 == "Gunting" {
		fmt.Println("Pemain 1 menang")
	} else if pemain1 == "Gunting" && pemain2 == "Kertas" {
		fmt.Println("Pemain 2 menang")
	} else if pemain1 == "Kertas" && pemain2 == "Batu" {
		fmt.Println("Pemain 1 menang")
	} else if pemain1 == "Batu" && pemain2 == "Kertas" {
		fmt.Println("Pemain 2 menang")
	} else if pemain1 == "Batu" && pemain2 == "Gunting" {
		fmt.Println("Pemain 1 menang")
	} else if pemain1 == "Gunting" && pemain2 == "Batu" {
		fmt.Println("Pemain 2 menang")
	} else if pemain1 == pemain2 {
		fmt.Println("Hasilnya draw")
	} else {
		fmt.Println("Ada pemain yang memilih diluar angka yang sudah diatur")
	}
}
