package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			m := k - i - j
			if m >= 1 && m <= n {
				res++
			}
		}
	}
	fmt.Println(res)
}
