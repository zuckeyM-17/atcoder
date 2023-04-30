package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	res := []string{}
	for i := 0; i <= 9; i++ {
		wari := math.Pow(2, float64(9-i))
		res = append(res, fmt.Sprint(n/int(wari)%2))
	}
	fmt.Println(strings.Join(res, ""))
}
