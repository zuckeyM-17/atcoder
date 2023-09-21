package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)

	a := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	grundy := make([]int, 100001)

	// Grundy 数を求める
	// 変数 grundy[i] : 石が i 個の山に対する Grundy 数
	// 変数 transit[i] : Grundy 数が i となるような遷移ができるかどうか

	for i := 0; i <= 100000; i++ {
		transit := make([]bool, 3)
		if i >= x {
			transit[grundy[i-x]] = true
		}
		if i >= y {
			transit[grundy[i-y]] = true
		}

		if !transit[0] {
			grundy[i] = 0
		} else if !transit[1] {
			grundy[i] = 1
		} else {
			grundy[i] = 2
		}
	}

	fmt.Fprintln(out, grundy)

	xorSum := 0
	for i := 1; i <= n; i++ {
		xorSum ^= grundy[a[i]]
	}

	if xorSum != 0 {
		fmt.Fprintln(out, "First")
	} else {
		fmt.Fprintln(out, "Second")
	}
}
