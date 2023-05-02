package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		if i == 0 {
			sum[i] = num
		} else {
			sum[i] = sum[i-1] + num
		}
	}

	var q int
	fmt.Scanf("%d", &q)

	var l, r int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d", &l, &r)
		s, t := sum[r]-sum[l-1], r-l+1
		if 2*s > t {
			fmt.Println("win")
		} else if 2*s == t {
			fmt.Println("draw")
		} else {
			fmt.Println("lose")
		}
	}
}
