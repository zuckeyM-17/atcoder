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

	a, b := make([]int, n), make([]int, n)

	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	dp := make([]int, n+1)

	dp[1] = 0
	for i := 1; i <= n-1; i++ {
		dp[a[i]] = max(dp[a[i]], dp[i]+100)
		dp[b[i]] = max(dp[b[i]], dp[i]+150)
	}

	fmt.Fprintln(out, dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
