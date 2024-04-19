package main

import "fmt"

func main(){
	// ellipsis operator
	var a5 = [...]int{
		1,
		-2,
		5,
		-10,
		25,
	}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("The length of a5 is %d\n", len(a5))
	fmt.Printf("The capacity of a5 is %d\n", cap(a5))

	a5[0] = 100
	fmt.Printf("%#v\n", a5)
} 