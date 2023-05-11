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

	var n, w int
	fmt.Fscan(in, &n, &w)

	ws, vs := make([]int, n+1), make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &ws[i], &vs[i])
	}

	fmt.Fprintln(out, knapsack(n, w, ws, vs))
}

func knapsack(n, w int, ws, vs []int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			if j < ws[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-ws[i]]+vs[i])
			}
		}
	}

	a := 0
	for i := 0; i <= w; i++ {
		a = max(a, dp[n][i])
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
