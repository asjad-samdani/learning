package main

import "fmt"

func main() {
	fmt.Println("Enter Numbers")
	var n int
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println(n, "Number is even")
	} else {
		fmt.Println(n, "Number is odd")

	}
}
