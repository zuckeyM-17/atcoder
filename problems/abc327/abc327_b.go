package main

import (
	"fmt"
	"math/big"
)

func main() {
	var B int64
	fmt.Scan(&B)

	for A := int64(1); A <= 20; A++ {
		bigA := big.NewInt(A)
		bigB := big.NewInt(B)
		if bigA.Exp(bigA, bigA, nil).Cmp(bigB) == 0 {
			fmt.Println(A)
			return
		}
	}
	fmt.Println(-1)
}
