package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int
		c int
		s string
	)
	fmt.Scan(&a)
	fmt.Scan(&b, &c)

	fmt.Scan(&s)

	fmt.Println(a+b+c, s)
}
