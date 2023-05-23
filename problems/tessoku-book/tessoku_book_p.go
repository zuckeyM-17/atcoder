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

	a, b := make([]int, n+2), make([]int, n+2)

	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	dp := make([]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = 1 << 60
	}
	dp[1] = 0
	for i := 1; i <= n-1; i++ {
		dp[i+1] = min(dp[i]+a[i+1], dp[i+1])
		dp[i+2] = min(dp[i]+b[i+2], dp[i+2])
	}

	fmt.Fprintln(out, dp[n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
