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

	var n int
	fmt.Fscan(in, &n)

	a, b := make([]int, n+1), make([]int, n+1)

	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = a[2]
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1]+a[i], dp[i-2]+b[i])
	}

	fmt.Fprintln(out, dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
