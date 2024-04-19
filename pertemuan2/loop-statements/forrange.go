package main

import "fmt"

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}