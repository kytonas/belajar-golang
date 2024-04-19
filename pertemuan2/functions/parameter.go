package main

import "fmt"

func helloParameter(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main(){
	helloParameter("Ksatria Faikar", "Nasywaan")
}