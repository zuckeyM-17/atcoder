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

	var N, W int
	fmt.Fscan(in, &N, &W)

	w, v := make([]int, N+1), make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &w[i], &v[i])
	}

	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, 100_000+1)
		for j := 0; j <= 100_000; j++ {
			dp[i][j] = 1<<32 - 1
		}
	}

	dp[0][0] = 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= 100_000; j++ {
			if j < v[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-v[i]]+w[i])
			}
		}

	}

	answer := 0
	for i := 0; i < 100000; i++ {
		if dp[N][i] <= W {
			answer = i
		}
	}
	fmt.Fprintln(out, answer)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
