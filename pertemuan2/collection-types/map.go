package main

import "fmt"

func main() {
	var employees map[string]string = map[string]string {
		"name": "Ksatria",
		"address": "Katapang",
	}

	fmt.Printf("%#v\n", employees)
	fmt.Println(employees["name"])
	fmt.Println(employees["address"])
	fmt.Printf("Length of employees is %d\n", len(employees))
}