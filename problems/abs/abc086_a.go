package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	times := a * b

	if times%2 > 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
