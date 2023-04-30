package main

import (
	"fmt"
)

func main() {
	var n, a int
	fmt.Scanf("%d", &n)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		as[i] = a
	}

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if as[i]+as[j]+as[k] == 1000 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}
