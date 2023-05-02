package main

import (
	"fmt"
)

func main() {
	var d, n int
	fmt.Scanf("%d", &d)
	fmt.Scanf("%d", &n)

	b := make([]int, d+1)

	for i := 0; i < n; i++ {
		var l, r int
		fmt.Scanf("%d %d", &l, &r)
		b[l-1]++
		b[r]--
	}

	var res int
	for _, v := range b[:len(b)-1] {
		res += v
		fmt.Println(res)
	}
}
