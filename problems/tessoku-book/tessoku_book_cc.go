package main

import (
	"fmt"
	"math"
)

func main() {
	var n string
	fmt.Scan(&n)

	res := 0
	for i, v := range n {
		if v == '0' {
			continue
		}

		res = res + int(math.Pow(2, float64(len(n)-1-i)))
	}

	fmt.Println(res)
}
