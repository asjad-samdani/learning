package main

import "fmt"

func main() {
	fmt.Println("Enter Two numbers")
	var x, y int

	fmt.Scanln(&x, &y)

	if x > y {
		fmt.Println(x, " is greater then ", y, "by", x-y)
	} else if x == y {
		fmt.Println(" Both value is equal ")

	} else {
		fmt.Println(y, " is greater then ", x, "by", y-x)

	}

}
