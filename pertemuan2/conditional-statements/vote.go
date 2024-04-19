package main

import "fmt"


func main(){
	var age int

	fmt.Println("Please input your age")
	fmt.Scanln(&age)

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return %d years\n", 18-age)
	} else if age == 18 {
		fmt.Println("This is your first vote")
	} else if age > 18 && age <= 100 {
		fmt.Println("Please vote, it's important")
	} else {
		fmt.Println("Invalid age!")
	}
}