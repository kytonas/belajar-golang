package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%v\n", nums)

	friends := [3]string{"John", "Paul", "George"}
	fmt.Printf("%v\n", friends)

	myBestFriend := friends[0:1]
	fmt.Println("My best friend is", myBestFriend)
	fmt.Printf("%v\n", myBestFriend)

	for index, value := range friends {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
}