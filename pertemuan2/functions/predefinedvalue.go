package main

import "fmt"

func getName() (firstName, lastName string) {
	return "Ksatria", "Nasywaan"
}

func main() {
	x, y := getName()
	fmt.Println(x, y)
}