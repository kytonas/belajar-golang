package main

import "fmt"

func main(){
	var language string

	fmt.Println("What's programming language you learn ?")
	fmt.Scanln(&language)

	switch language {
	case "Python":
		fmt.Println("You are learning Python!")
	case "Go", "Golang":
		fmt.Println("Good, go for Go")
	case "PHP":
		fmt.Println("PHP is good for beginner")
	case "C++", "C#":
		fmt.Println("Sheesh bro, you just like computer")
	case "Assembly":
		fmt.Println("Mati Gue")
	default:
		fmt.Println("Any Other programming language is a good start")
	}
}