package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)
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

	var l, r int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d", &l, &r)
		fmt.Println(sum[r] - sum[l-1])
	}
}
